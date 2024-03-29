package web

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gorilla/securecookie"
	"github.com/imrancluster/th-common-payment/config"
	"github.com/imrancluster/th-common-payment/conn"
	"github.com/imrancluster/th-common-payment/models"
	"golang.org/x/crypto/bcrypt"
)

// cookie handling
var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

// UserInfo ..
type UserInfo struct {
	ID       int
	Username string
	Email    string
}

// GetUserInfoFromSession ..
func GetUserInfoFromSession(request *http.Request) UserInfo {

	var userInfo UserInfo

	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {

			ID, err := strconv.Atoi(cookieValue["ID"])
			if err != nil {
				fmt.Println(err)
			}

			userInfo = UserInfo{
				ID:       ID,
				Username: cookieValue["username"],
				Email:    cookieValue["email"],
			}
		}
	}

	return userInfo
}

func setSession(userInfo models.User, response http.ResponseWriter) {
	value := map[string]string{
		"ID":       fmt.Sprint(userInfo.ID),
		"username": userInfo.Username,
		"email":    userInfo.Email,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

// UserWeb ..
type UserWeb struct{}

type userData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SignUp ..
func (u *UserWeb) SignUp(w http.ResponseWriter, r *http.Request) {

	config.TPL.ExecuteTemplate(w, "sign-up.gohtml", nil)
}

// SignUpProcess ..
func (u *UserWeb) SignUpProcess(w http.ResponseWriter, r *http.Request) {
	var user userData
	var msg Message
	msg.Errors = make(map[string]string)

	re := regexp.MustCompile(".+@.+\\..+")
	matched := re.Match([]byte(r.FormValue("email")))
	if matched == false {
		msg.Errors["Email"] = "Please enter a valid email address"
	}

	if strings.TrimSpace(r.FormValue("username")) == "" {
		msg.Errors["Username"] = "Username should not empty"
	}

	if len(strings.TrimSpace(r.FormValue("password"))) < 6 {
		msg.Errors["Password"] = "Password should be minimum 6 digit"
	}

	if len(msg.Errors) != 0 {
		msg.Errors["EmailValue"] = r.FormValue("email")
		msg.Errors["UsernameValue"] = r.FormValue("username")
		config.TPL.ExecuteTemplate(w, "sign-up.gohtml", msg)
		return
	}

	user.Username = r.FormValue("username")
	user.Password = r.FormValue("password")
	user.Email = r.FormValue("email")

	passHash, errHash := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if errHash != nil {
		http.Error(w, errHash.Error(), http.StatusForbidden)
		return
	}

	readerUser := models.User{
		Username: user.Username,
		Email:    user.Email,
		Password: string(passHash),
	}

	db := conn.PostgresDB()
	err := db.Create(&readerUser).Error
	if err != nil {
		fmt.Println("Error: ", err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// SignIn ..
func (u *UserWeb) SignIn(w http.ResponseWriter, r *http.Request) {

	config.TPL.ExecuteTemplate(w, "sign-in.gohtml", nil)
}

// SignInProcess ..
func (u *UserWeb) SignInProcess(w http.ResponseWriter, r *http.Request) {

	var msg Message
	msg.Errors = make(map[string]string)

	var user models.User
	db := conn.PostgresDB()

	if err := db.Where("username = ?", r.FormValue("username")).First(&user).Error; err != nil {
		msg.Errors["Login"] = "Username and/or password do not match"
	}

	// does the entered password match the stored password?
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.FormValue("password")))
	if err != nil {
		msg.Errors["Login"] = "Username and/or password do not match"
	}

	if len(msg.Errors) != 0 {
		config.TPL.ExecuteTemplate(w, "sign-in.gohtml", msg)
		return
	}

	// .. check credentials ..
	setSession(user, w)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// SignOut ..
func (u *UserWeb) SignOut(w http.ResponseWriter, r *http.Request) {
	clearSession(w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Profile ..
func (u *UserWeb) Profile(w http.ResponseWriter, r *http.Request) {

	userInfo := GetUserInfoFromSession(r)

	config.TPL.ExecuteTemplate(w, "profile.gohtml", userInfo)
}

// NewWebUser ..
func NewWebUser() *UserWeb {
	return &UserWeb{}
}
