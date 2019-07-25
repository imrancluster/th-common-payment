package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/imrancluster/th-common-payment/handlers"
)

func userWebRoutes() http.Handler {
	route := chi.NewRouter()
	route.Group(func(r chi.Router) {
		r.Get("/signup", handlers.NewWebUser().SignUp)
		r.Post("/signup/process", handlers.NewWebUser().SignUpProcess)
		r.Get("/signin", handlers.NewWebUser().SignIn)
		r.Post("/signin/process", handlers.NewWebUser().SignInProcess)
	})

	return route
}

func userRoutes() http.Handler {
	route := chi.NewRouter()
	route.Group(func(r chi.Router) {
		r.Post("/", handlers.NewUserAPI().CreateUser)
		r.Get("/{user_id:[0-9]+}", handlers.NewUserAPI().GetUser)
	})
	return route
}
