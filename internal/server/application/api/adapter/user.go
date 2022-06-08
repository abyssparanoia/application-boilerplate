package adapter

import (
	"time"

	"github.com/abyssparanoia/application-boilerplate/internal/server/application/api/generated/openapi"
	"github.com/abyssparanoia/application-boilerplate/internal/server/domain/model"
)

func NewUser(m *model.User) openapi.User {
	if m == nil {
		return openapi.User{}
	}
	return openapi.User{
		Id:                  m.ID,
		DisplayName:         m.DisplayName,
		IconImagePath:       m.IconImagePath,
		BackgroundImagePath: m.BackgroundImagePath,
		Profile:             m.Profile,
		Email:               m.Email,
		CreatedAt:           m.CreatedAt.Format(time.RFC3339),
		UpdatedAt:           m.UpdatedAt.Format(time.RFC3339),
	}
}

func NewUsers(ms []*model.User) []openapi.User {
	dsts := make([]openapi.User, len(ms))
	for idx, m := range ms {
		dsts[idx] = NewUser(m)
	}
	return dsts
}
