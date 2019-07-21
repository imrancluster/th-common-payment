package admin

import (
	"errors"
	"net/http"
	"time"

	"../config"
)

type User struct {
	Username  string
	Password  string
	Email     string
	CreatedOn time.Time
	LastLogin time.Time
}

func InsertUser(r *http.Request) (User, error) {

	// Get Form Values
	u := User{}
	u.Username = r.FormValue("username")
	u.Password = r.FormValue("password")
	u.Email = r.FormValue("password")
	u.CreatedOn = time.Now()

	// validate form
	// validate form values
	if u.Username == "" || u.Password == "" || u.Email == "" {
		return u, errors.New("400. Bad request. All fields must be complete.")
	}

	// insert values
	_, err := config.DB.Exec("INSERT INTO users (username, password, email, created_on) VALUES ($1, $2, $3, $4)", u.Username, u.Password, u.Email, u.CreatedOn)
	if err != nil {
		return u, errors.New("500. Internal Server Error." + err.Error())
	}

	return u, nil
}
