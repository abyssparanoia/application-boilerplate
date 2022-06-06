package api

import (
	"net/http"

	"github.com/abyssparanoia/application-boilerplate/internal/pkg/accesscontrol"
	"github.com/abyssparanoia/application-boilerplate/internal/server/application/api/api_handlers"
	"github.com/go-chi/chi"
)

// Routing ... define routing
func Routing(r chi.Router, d Dependency) {

	// access control
	r.Use(accesscontrol.Handle)

	// request log
	r.Use(d.httpMiddleware.Handle)

	// need to authenticate for production
	r.Route("/v1", func(r chi.Router) {
		r.With(d.authenticationMiddleware.Handle).Route("/users", func(r chi.Router) {
			r.Get("/{userID}", d.userHandler.Get)
		})
	})

	// Ping
	r.Get("/ping", api_handlers.Ping)
	r.Get("/", api_handlers.Ping)

	http.Handle("/", r)
}
