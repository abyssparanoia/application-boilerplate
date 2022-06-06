package repository

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/abyssparanoia/application-boilerplate/internal/dbmodels/defaultdb"
	"github.com/abyssparanoia/application-boilerplate/internal/pkg/gluesqlboiler"
	"github.com/abyssparanoia/application-boilerplate/internal/server/domain/model"
	"github.com/abyssparanoia/application-boilerplate/internal/server/domain/repository"
	"github.com/abyssparanoia/application-boilerplate/internal/server/infrastructure/internal/entity"
	"github.com/abyssparanoia/application-boilerplate/internal/server/pkg/server_error"
)

type user struct {
}

func (r *user) Get(ctx context.Context, userID string) (*model.User, error) {

	dbUser, err := defaultdb.Users(
		defaultdb.UserWhere.ID.EQ(userID),
	).One(ctx, gluesqlboiler.GetContextExecutor(ctx))

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, server_error.NotFoundErr.Errorf("User %s is not found", userID)
		}
		return nil, server_error.InternalError.Wrap(err)
	}

	user := entity.User{User: *dbUser}
	return user.OutputModel(), nil
}

func (r *user) Create(
	ctx context.Context,
	user *model.User,
) (*model.User, error) {
	dbUser := entity.NewUserFromModel(user)

	err := dbUser.Insert(ctx, gluesqlboiler.GetContextExecutor(ctx), boil.Infer())
	if err != nil {
		return nil, server_error.InternalError.Wrap(err)
	}

	return dbUser.OutputModel(), nil
}

// NewUser ... get user repository
func NewUser() repository.User {
	return &user{}
}
