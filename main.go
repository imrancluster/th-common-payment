// entry point of application

package main

// Need to add personal git package here, like controllers, models etc
// "github.com/imrancluster/th-common-payment/config"
// "github.com/imrancluster/th-common-payment/admin"
import (
	"net/http"

	"./admin"
	"./config"
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

	config.CLog("VISIT_INDEX_PAGE", "01799997163", map[string]interface{}{"data1": "test1", "data2": "test2"}).Info("Visit Index Page")
	config.CLog("INDEX_WARNING", "01799997163", map[string]interface{}{"data1": "t1", "data2": "t2"}).Warning("Index warning")
	config.CLog("INDEX_ERROR", "01799997163", map[string]interface{}{"data1": "t1", "data2": "t2"}).Error("Index Error")

	config.TPL.ExecuteTemplate(w, "index.gohtml", nil)
}
