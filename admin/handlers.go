package admin

import (
	"net/http"

	"../config"
	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	config.TPL.ExecuteTemplate(w, "admin-index.gohtml", nil)
}
