package admin

import (
	"net/http"

	"github.com/imrancluster/th-common-payment/config"
	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	config.TPL.ExecuteTemplate(w, "index.html", nil)
}
