package usecase

import (
	"context"

	"github.com/abyssparanoia/application-boilerplate/internal/server/domain/model"
	"github.com/abyssparanoia/application-boilerplate/internal/server/domain/repository"
)

type authentication struct {
	rAuthentication repository.Authentication
}

func NewAuthentication(
	rAuthentication repository.Authentication,
) Authentication {
	return &authentication{
		rAuthentication,
	}
}

func (u *authentication) Verify(ctx context.Context, idToken string) (*model.AuthUserInfo, error) {
	return u.rAuthentication.VerifyIDToken(ctx, idToken)
}
