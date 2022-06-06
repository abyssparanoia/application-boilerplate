package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/abyssparanoia/application-boilerplate/internal/pkg/renderer"
	"github.com/abyssparanoia/application-boilerplate/internal/server/domain/model"
	"github.com/abyssparanoia/application-boilerplate/internal/server/pkg/server_error"
	"github.com/abyssparanoia/application-boilerplate/internal/server/usecase"
)

type Authentication struct {
	uAuthentication usecase.Authentication
}

func NewAuthentication(
	uAuthentication usecase.Authentication,
) *Authentication {
	return &Authentication{
		uAuthentication,
	}
}

const (
	headerPrefix           string = "Bearer"
	authUserInfoContextKey string = "authUserInfo"
)

func getTokenByAuthHeader(ah string) string {
	pLen := len(headerPrefix)
	if len(ah) > pLen && strings.ToUpper(ah[0:pLen]) == headerPrefix {
		return ah[pLen+1:]
	}
	return ""
}

func (m *Authentication) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ah := r.Header.Get("Authorization")
		token := getTokenByAuthHeader(ah)
		if token == "" {
			renderer.HandleError(ctx, w, server_error.UnauthorizedErr.Errorf("No Authorization Header"))
			return
		}

		authUserInfo, err := m.uAuthentication.Verify(ctx, token)
		if err != nil {
			renderer.HandleError(ctx, w, server_error.UnauthorizedErr.Wrap(err))
			return
		}
		ctx = SetAuthUserInfo(ctx, authUserInfo)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func SetAuthUserInfo(ctx context.Context, authUserInfo *model.AuthUserInfo) context.Context {
	return context.WithValue(ctx, authUserInfoContextKey, *authUserInfo)
}

func GetAuthUserInfo(ctx context.Context) *model.AuthUserInfo {
	userInfo := ctx.Value(authUserInfoContextKey).(model.AuthUserInfo)
	return &userInfo
}
