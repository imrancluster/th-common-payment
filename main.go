// entry point of application

package main

// Need to add personal git package here, like controllers, models etc
import (
	"html/template"
	"net/http"

	"github.com/imrancluster/th-common-payment/controllers"
	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	// initiate httprouter
	route := httprouter.New()

	// initate front page controller
	frontController := controllers.NewFrontController(tpl)
	adminController := controllers.NewAdminController(tpl)

	// all routers
	route.GET("/", frontController.HomePage)
	route.GET("/admin", adminController.Index)

	// start server
	http.ListenAndServe("localhost:8080", route)

}
