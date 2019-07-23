package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/imrancluster/th-common-payment/models"
	"github.com/imrancluster/th-common-payment/repository"
	userRepoImpl "github.com/imrancluster/th-common-payment/repository/user"
	"golang.org/x/crypto/bcrypt"
)

// UserAPI ..
type UserAPI struct {
	repo repository.UserRepo
}

// CreateUser ..
func (u *UserAPI) CreateUser(w http.ResponseWriter, r *http.Request) {
	type userData struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

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

// NewUserAPI ..
func NewUserAPI() *UserAPI {
	return &UserAPI{
		repo: userRepoImpl.NewUser(),
	}
}
