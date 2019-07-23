package admin

import (
	"net/http"

	"github.com/imrancluster/th-common-payment/config"
	"github.com/julienschmidt/httprouter"
)

// Index is a route function for admin front page
func Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	config.TPL.ExecuteTemplate(w, "admin-index.gohtml", nil)
}

// AllUsers ..
func AllUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	config.TPL.ExecuteTemplate(w, "all-users.gohtml", nil)
}

// AddUser ..
func AddUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	config.TPL.ExecuteTemplate(w, "add-user.gohtml", nil)
}

// UserProcess ..
func UserProcess(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	_, err := InsertUser(r)
	if err != nil {
		config.CLog("USER_CREATE_NOT_ACCEPTABLE", "", map[string]interface{}{"error": http.StatusNotAcceptable}).Info("User submit data not acceptable")
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
		return
	}

	// Logging
	config.CLog("NEW_USER_CREATED", "", map[string]interface{}{"response": "r"}).Info("User has been created")

	http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
}
