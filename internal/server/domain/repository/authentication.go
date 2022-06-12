package repository

import (
	"context"

	"github.com/abyssparanoia/application-boilerplate/internal/server/domain/model"
)

//go:generate mockgen -source authentication.go -destination mock/mock_authentication.go
type Authentication interface {
	VerifyIDToken(ctx context.Context, idToken string) (*model.AuthUserInfo, error)
	StoreClaims(ctx context.Context, userID string, claims *model.Claims) error
}
