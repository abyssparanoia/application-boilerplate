package api_handlers

import (
	"net/http"

	"github.com/abyssparanoia/application-boilerplate/internal/pkg/parameter"
	"github.com/abyssparanoia/application-boilerplate/internal/pkg/renderer"
	"github.com/abyssparanoia/application-boilerplate/internal/server/application/api/adapter"
	"github.com/abyssparanoia/application-boilerplate/internal/server/application/api/generated/openapi"
	"github.com/abyssparanoia/application-boilerplate/internal/server/usecase"
)

// UserHandler ... user handler
type UserHandler struct {
	userUsecase usecase.User
}

// Get ... get user
func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user, err := h.userUsecase.Get(ctx, parameter.GetURL(r, "userID"))
	if err != nil {
		renderer.HandleError(ctx, w, err)
		return
	}

	renderer.JSON(ctx, w, http.StatusOK, &openapi.GetUserResponse{
		User: adapter.NewUser(user),
	})
}

// NewUserHandler ... get user handler
func NewUserHandler(userUsecase usecase.User) *UserHandler {
	return &UserHandler{userUsecase}
}
