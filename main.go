// entry point of application

package main

// Need to add personal git package here, like controllers, models etc
import (
	"net/http"

	"github.com/imrancluster/th-common-payment/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {

	// initiate httprouter
	route := httprouter.New()

	// initate front page controller
	frontController := controllers.NewFrontController()

	// all routers
	route.GET("/", frontController.HomePage)

	// start server
	http.ListenAndServe("localhost:8080", route)

}
