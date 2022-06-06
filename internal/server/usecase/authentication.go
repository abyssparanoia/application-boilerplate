package usecase

import (
	"context"

	"github.com/abyssparanoia/application-boilerplate/internal/server/domain/model"
)

type Authentication interface {
	Verify(ctx context.Context, idToken string) (*model.AuthUserInfo, error)
}
