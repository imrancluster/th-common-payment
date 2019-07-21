package main

import (
	"net/http"

	"./admin"
	"./config"
	"github.com/julienschmidt/httprouter"
)

func main() {

	route := httprouter.New()
	route.ServeFiles("/assets/*filepath", http.Dir("public"))
	route.GET("/", index)
	route.GET("/admin", admin.Index)
	route.GET("/admin/users", admin.AllUsers)
	route.GET("/admin/users/add", admin.AddUser)
	route.POST("/admin/users/process", admin.UserProcess)

	http.ListenAndServe("localhost:8080", route)
}

func index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	config.TPL.ExecuteTemplate(w, "index.gohtml", nil)
}
