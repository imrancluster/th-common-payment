package admin

import (
	"net/http"

	"../config"
	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	config.TPL.ExecuteTemplate(w, "admin-index.gohtml", nil)
}

func AllUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	config.TPL.ExecuteTemplate(w, "all-users.gohtml", nil)
}

func AddUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	config.TPL.ExecuteTemplate(w, "add-user.gohtml", nil)
}

func UserProcess(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	_, err := InsertUser(r)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
		return
	}

	http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
}
