// entry point of application

package main

// Need to add personal git package here, like controllers, models etc
import (
	"log"
	"net/http"

	"github.com/imrancluster/th-common-payment/admin"
	"github.com/imrancluster/th-common-payment/config"
	"github.com/julienschmidt/httprouter"
)

func main() {

	// initiate httprouter
	route := httprouter.New()

	// all routes
	route.GET("/", index)
	route.GET("/admin", admin.Index)

	// start server
	http.ListenAndServe("localhost:8080", route)

}

func index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Println("Index page")
	config.TPL.ExecuteTemplate(w, "index.gohtml", nil)
}
