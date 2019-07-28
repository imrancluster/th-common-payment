package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/imrancluster/th-common-payment/handlers/api"
	"github.com/imrancluster/th-common-payment/handlers/web"
)

func userWebRoutes() http.Handler {
	route := chi.NewRouter()
	route.Group(func(r chi.Router) {
		r.Get("/signup", web.NewWebUser().SignUp)
		r.Post("/signup/process", web.NewWebUser().SignUpProcess)
		r.Get("/signin", web.NewWebUser().SignIn)
		r.Post("/signin/process", web.NewWebUser().SignInProcess)
		r.Post("/signout", web.NewWebUser().SignOut)
		r.Get("/profile", web.NewWebUser().Profile)
	})

	return route
}

func userRoutes() http.Handler {
	route := chi.NewRouter()
	route.Group(func(r chi.Router) {
		r.Post("/", api.NewUserAPI().CreateUser)
		r.Get("/{user_id:[0-9]+}", api.NewUserAPI().GetUser)
	})
	return route
}
