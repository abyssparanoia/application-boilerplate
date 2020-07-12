package server

import (
	"net/http"

	"github.com/abyssparanoia/application-boilerplate/internal/pkg/accesscontrol"
	"github.com/abyssparanoia/application-boilerplate/internal/server/handler"
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
		r.With(d.gluefirebaseauth.Handle).Route("/users", func(r chi.Router) {
			//r.Post("/", d.UserHandler.Create)
			r.Get("/{userID}", d.UserHandler.Get)
		})
	})

	// Ping
	r.Get("/ping", handler.Ping)
	r.Get("/", handler.Ping)

	http.Handle("/", r)
}
