package repository

import (
	"context"

	"github.com/abyssparanoia/application-boilerplate/internal/server/domain/model"
)

//go:generate mockgen -source user.go -destination mock/mock_user.go
type User interface {
	Get(
		ctx context.Context,
		userID string,
	) (*model.User, error)
	Create(
		ctx context.Context,
		user *model.User,
	) (*model.User, error)
}
