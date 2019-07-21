// entry point of application

package main

// Need to add personal git package here, like controllers, models etc
import (
	"html/template"
	"net/http"

	"github.com/imrancluster/th-common-payment/admin"
	"github.com/imrancluster/th-common-payment/config"
	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	// initiate httprouter
	route := httprouter.New()

	// all routes
	route.GET("/", index)
	route.GET("/admin", admin.Index)

	// start server
	http.ListenAndServe("localhost:8080", route)

}

func index(w http.ResponseWriter, r *http.Request) {
	config.TPL.ExecuteTemplate(w, "index.html", nil)
}
