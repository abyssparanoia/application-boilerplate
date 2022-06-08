package repository

import (
	"context"

	"firebase.google.com/go/auth"
	"github.com/abyssparanoia/application-boilerplate/internal/pkg/glueerr"
	"github.com/abyssparanoia/application-boilerplate/internal/server/domain/model"
	"github.com/abyssparanoia/application-boilerplate/internal/server/domain/repository"
	"github.com/abyssparanoia/application-boilerplate/internal/server/infrastructure/internal/dto"
)

type authentication struct {
	cli *auth.Client
}

func NewAuthentication(
	cli *auth.Client,
) repository.Authentication {
	return &authentication{
		cli,
	}
}

func (r *authentication) VerifyIDToken(ctx context.Context, idToken string) (*model.AuthUserInfo, error) {
	t, err := r.cli.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, glueerr.UnauthorizedErr.Wrap(err)
	}
	userInfo := &model.AuthUserInfo{}
	userInfo.UserID = t.UID
	userInfo.Claims = dto.SetMapToClaims(t.Claims)
	return userInfo, nil
}

func (r *authentication) StoreClaims(ctx context.Context, userID string, claims *model.Claims) error {
	if err := r.cli.SetCustomUserClaims(ctx, userID, dto.ClaimsToMap(claims)); err != nil {
		return glueerr.InternalErr.Wrap(err)
	}
	return nil
}
