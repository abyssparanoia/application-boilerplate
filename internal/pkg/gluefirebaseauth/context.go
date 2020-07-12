package gluefirebaseauth

import "context"

type contextKey string

const (
	userIDContextKey     contextKey = "gluefirebaseauth:user_id"
	authHeaderContextKey contextKey = "gluefirebaseauth:auth_header"

	claimsContextKey contextKey = "gluefirebaseauth:claims"
)

func getAuthHeader(ctx context.Context) string {
	return ctx.Value(authHeaderContextKey).(string)
}

// GetUserID ... get gluefirebaseauthUID from context
func GetUserID(ctx context.Context) string {
	return ctx.Value(userIDContextKey).(string)
}

// GetClaims ... get gluefirebaseauthJWTClaims from context
func GetClaims(ctx context.Context) Claims {
	return ctx.Value(claimsContextKey).(Claims)
}

func setAuthHeader(ctx context.Context, ah string) context.Context {
	return context.WithValue(ctx, authHeaderContextKey, ah)
}

func setUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDContextKey, userID)
}

func setClaims(ctx context.Context, claims *Claims) context.Context {
	return context.WithValue(ctx, claimsContextKey, claims)
}
