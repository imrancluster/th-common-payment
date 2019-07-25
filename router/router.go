package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/imrancluster/th-common-payment/config"
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

	config.TPL.ExecuteTemplate(w, "index.gohtml", nil)
}

// Router defines all routers
func Router() *chi.Mux {

	router.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir("public"))))

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
