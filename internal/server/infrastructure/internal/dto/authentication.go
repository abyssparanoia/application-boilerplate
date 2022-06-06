package dto

import (
	"github.com/abyssparanoia/application-boilerplate/internal/server/domain/model"
)

func SetMapToClaims(cmap map[string]interface{}) *model.Claims {
	return &model.Claims{}
}

func ClaimsToMap(claims *model.Claims) map[string]interface{} {
	cmap := map[string]interface{}{}
	return cmap
}
