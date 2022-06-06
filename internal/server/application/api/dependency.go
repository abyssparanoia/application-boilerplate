package api

import (
	"github.com/abyssparanoia/application-boilerplate/internal/pkg/httpheader"
	"github.com/abyssparanoia/application-boilerplate/internal/pkg/middleware/requestlog"
	"github.com/abyssparanoia/application-boilerplate/internal/server/application/api/api_handlers"
	"github.com/abyssparanoia/application-boilerplate/internal/server/application/api/middlewares"
	"github.com/abyssparanoia/application-boilerplate/internal/server/application/shared"
	"go.uber.org/zap"
)

// Dependency ... dependency
type Dependency struct {
	httpMiddleware           *requestlog.HTTPMiddleware
	dummyHTTPHeader          *httpheader.Middleware
	hTTPHeader               *httpheader.Middleware
	authenticationMiddleware *middlewares.Authentication

	userHandler *api_handlers.UserHandler
}

// Inject ... indect dependency
func (d *Dependency) Inject(e *shared.Environment, logger *zap.Logger) {

	sharedDependency := &shared.Dependency{}
	sharedDependency.Inject(e, logger)

	dhh := httpheader.NewDummy()
	hh := httpheader.New()

	// Middleware
	d.httpMiddleware = requestlog.New(logger)
	d.authenticationMiddleware = middlewares.NewAuthentication(sharedDependency.AuthenticationUsecase)
	d.dummyHTTPHeader = httpheader.NewMiddleware(dhh)
	d.hTTPHeader = httpheader.NewMiddleware(hh)

	// Handler
	d.userHandler = api_handlers.NewUserHandler(sharedDependency.UserUsecase)
}
