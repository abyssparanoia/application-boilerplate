package testdata

import (
	"time"

	"github.com/abyssparanoia/application-boilerplate/internal/server/domain/model"
	"github.com/bxcodec/faker"
)

// NewDomainModel :
func NewDomainModel() *struct {
	User         *model.User
	AuthUserInfo *model.AuthUserInfo
} {

	user := &model.User{}
	if err := faker.FakeData(user); err != nil {
		panic(err)
	}

	user.CreatedAt = time.Time{}
	user.UpdatedAt = time.Time{}

	authUserInfo := &model.AuthUserInfo{}
	authUserInfo.UserID = user.ID
	authUserInfo.Claims = &model.Claims{}

	return &struct {
		User         *model.User
		AuthUserInfo *model.AuthUserInfo
	}{
		User:         user,
		AuthUserInfo: authUserInfo,
	}
}
