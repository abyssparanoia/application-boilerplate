package server

import (
	"github.com/abyssparanoia/application-boilerplate/internal/pkg/gluefirebaseauth"
	"github.com/abyssparanoia/application-boilerplate/internal/pkg/gluemysql"
	"github.com/abyssparanoia/application-boilerplate/internal/pkg/httpheader"
	"github.com/abyssparanoia/application-boilerplate/internal/pkg/middleware/requestlog"
	"github.com/abyssparanoia/application-boilerplate/internal/server/handler/api"
	"github.com/abyssparanoia/application-boilerplate/internal/server/infrastructure/repository"
	"github.com/abyssparanoia/application-boilerplate/internal/server/usecase"
	"github.com/volatiletech/sqlboiler/boil"
	"go.uber.org/zap"
)

// Dependency ... dependency
type Dependency struct {
	httpMiddleware   *requestlog.HTTPMiddleware
	gluefirebaseauth *gluefirebaseauth.Middleware
	DummyHTTPHeader  *httpheader.Middleware
	HTTPHeader       *httpheader.Middleware
	UserHandler      *api.UserHandler
}

// Inject ... indect dependency
func (d *Dependency) Inject(e *environment, logger *zap.Logger) {

	var firebaseauth gluefirebaseauth.Firebaseauth

	authCli := gluefirebaseauth.NewClient(e.ProjectID)
	// fCli := gluefirestore.NewClient(e.ProjectID)

	// pkg
	_ = gluemysql.NewClient(e.DBHost, e.DBUser, e.DBPassword, e.DBDatabase)

	if e.Envrionment == "local" {
		firebaseauth = gluefirebaseauth.NewDebug(authCli)
		boil.DebugMode = true
	} else {
		firebaseauth = gluefirebaseauth.New(authCli)
	}

	// Repository
	uRepo := repository.NewUser()

	// Service
	dhh := httpheader.NewDummy()
	hh := httpheader.New()
	uSvc := usecase.NewUser(uRepo)

	// Middleware
	d.httpMiddleware = requestlog.New(logger)

	d.gluefirebaseauth = gluefirebaseauth.NewMiddleware(firebaseauth)
	d.DummyHTTPHeader = httpheader.NewMiddleware(dhh)
	d.HTTPHeader = httpheader.NewMiddleware(hh)

	// Handler
	d.UserHandler = api.NewUserHandler(uSvc)
}
