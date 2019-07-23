package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/imrancluster/th-common-payment/handlers"
)

func userRoutes() http.Handler {
	route := chi.NewRouter()
	route.Group(func(r chi.Router) {
		r.Post("/", handlers.NewUserAPI().CreateUser)
		r.Get("/{user_id:[0-9]+}", handlers.NewUserAPI().GetUser)
	})
	return route
}
