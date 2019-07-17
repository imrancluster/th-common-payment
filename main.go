// entry point of application

package main

// Need to add personal git package here, like controllers, models etc
import (
	"github.com/julienschmidt/httprouter"
)

func main() {

	// initiate httprouter
	route := httprouter.New()

	// initate front page controller
	frontController := controllers.NewFrontController()

	// all routers

}
