package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/imrancluster/th-common-payment/config"
	"github.com/imrancluster/th-common-payment/conn"
	"github.com/imrancluster/th-common-payment/models"
	"github.com/imrancluster/th-common-payment/repository"
	userRepoImpl "github.com/imrancluster/th-common-payment/repository/user"
	"golang.org/x/crypto/bcrypt"
)

// UserAPI ..
type UserAPI struct {
	repo repository.UserRepo
}

type UserWeb struct{}

type userData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CreateUser ..
func (u *UserAPI) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user userData

	body := json.NewDecoder(r.Body)
	if err := body.Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	passHash, errHash := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if errHash != nil {
		respondWithError(w, http.StatusBadRequest, errHash.Error())
		return
	}

	readerUser := models.User{
		Username: user.Username,
		Email:    user.Email,
		Password: string(passHash),
	}

	payload, errCrUser := u.repo.CreateUser(r.Context(), &readerUser)
	if errCrUser != nil {
		respondWithError(w, http.StatusConflict, errCrUser.Error())
		return
	}

	respondwithJSON(w, http.StatusCreated, payload)
}

// GetUser ..
func (u *UserAPI) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "user_id"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	getUser, errGet := u.repo.GetUser(r.Context(), id)

	if errGet != nil {
		respondWithError(w, http.StatusNotFound, fmt.Sprintf("%s", errGet))
		return
	}

	respondwithJSON(w, http.StatusOK, getUser)
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
		respondWithError(w, http.StatusBadRequest, errHash.Error())
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

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// NewUserAPI ..
func NewUserAPI() *UserAPI {
	return &UserAPI{
		repo: userRepoImpl.NewUser(),
	}
}

// NewWebUser ..
func NewWebUser() *UserWeb {
	return &UserWeb{}
}
