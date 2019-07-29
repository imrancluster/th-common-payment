package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/imrancluster/th-common-payment/config"
	"github.com/imrancluster/th-common-payment/handlers/web"
)

var router *chi.Mux

// route.ServeFiles("/assets/*filepath", http.Dir("public"))

func init() {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
}

// SystemRoot handller
func SystemRoot(w http.ResponseWriter, r *http.Request) {

	userInfo := web.GetUserInfoFromSession(r)

	config.TPL.ExecuteTemplate(w, "index.gohtml", userInfo)
}

// Router defines all routers
func Router() *chi.Mux {

	router.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir("public"))))
	router.Handle("/favicon.ico", http.NotFoundHandler())

	// Routes for Web
	router.Get("/", SystemRoot)
	router.Route("/", func(r chi.Router) {
		r.Mount("/user", userWebRoutes())
	})

	router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/users", userRoutes())
	})

	return router
}
