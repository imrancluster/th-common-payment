package web

import (
	"fmt"
	"net/http"

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

func GetUserName(request *http.Request) (userName string) {
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}

func setSession(userName string, response http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
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

// SignIn
func (u *UserWeb) SignIn(w http.ResponseWriter, r *http.Request) {

	config.TPL.ExecuteTemplate(w, "sign-in.gohtml", nil)
}

func (u *UserWeb) SignInProcess(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	password := r.FormValue("password")

	var user models.User
	db := conn.PostgresDB()

	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		http.Error(w, "Username and/or password do not match", http.StatusForbidden)
		return
	}

	// does the entered password match the stored password?
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		http.Error(w, "Username and/or password do not match", http.StatusForbidden)
		return
	}

	// .. check credentials ..
	setSession(user.Username, w)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// SignOut ..
func (u *UserWeb) SignOut(w http.ResponseWriter, r *http.Request) {
	clearSession(w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (u *UserWeb) Profile(w http.ResponseWriter, r *http.Request) {

	type UserInfo struct{ Username string }
	userInfo := UserInfo{Username: GetUserName(r)}

	config.TPL.ExecuteTemplate(w, "profile.gohtml", userInfo)
}

// NewWebUser ..
func NewWebUser() *UserWeb {
	return &UserWeb{}
}
