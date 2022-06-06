package shared

import (
	"github.com/abyssparanoia/application-boilerplate/internal/pkg/gluefirebaseauth"
	"github.com/abyssparanoia/application-boilerplate/internal/pkg/gluemysql"
	"github.com/abyssparanoia/application-boilerplate/internal/server/infrastructure/repository"
	"github.com/abyssparanoia/application-boilerplate/internal/server/usecase"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.uber.org/zap"
)

// Dependency ... dependency
type Dependency struct {
	UserUsecase           usecase.User
	AuthenticationUsecase usecase.Authentication
}

// Inject ... indect dependency
func (d *Dependency) Inject(e *Environment, logger *zap.Logger) {
	firebaseAuthCli := gluefirebaseauth.NewClient(e.ProjectID)

	// pkg
	_ = gluemysql.NewClient(e.DBHost, e.DBUser, e.DBPassword, e.DBDatabase)

	if e.Envrionment == "local" {
		boil.DebugMode = true
	} else {
	}

	// Repository
	rUser := repository.NewUser()
	rAuthentication := repository.NewAuthentication(firebaseAuthCli)

	// Service
	d.UserUsecase = usecase.NewUser(rUser)
	d.AuthenticationUsecase = usecase.NewAuthentication(rAuthentication)

}
