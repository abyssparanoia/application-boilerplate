package api_handlers

import (
	"net/http"

	"github.com/abyssparanoia/application-boilerplate/internal/pkg/parameter"
	"github.com/abyssparanoia/application-boilerplate/internal/pkg/renderer"
	"github.com/abyssparanoia/application-boilerplate/internal/server/domain/model"
	"github.com/abyssparanoia/application-boilerplate/internal/server/usecase"
	validator "github.com/go-playground/validator"
)

// UserHandler ... user handler
type UserHandler struct {
	userUsecase usecase.User
}

type userHandlerGetRequestParam struct {
	UserID string `validate:"required"`
}

type userHandlerGetResponse struct {
	User *model.User `json:"user"`
}

// Get ... get user
func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var param userHandlerGetRequestParam
	param.UserID = parameter.GetURL(r, "userID")

	v := validator.New()
	if err := v.Struct(param); err != nil {
		renderer.HandleError(ctx, w, err)
		return
	}

	user, err := h.userUsecase.Get(ctx, param.UserID)
	if err != nil {
		renderer.HandleError(ctx, w, err)
		return
	}

	renderer.JSON(ctx, w, http.StatusOK, userHandlerGetResponse{User: user})
}

// NewUserHandler ... get user handler
func NewUserHandler(userUsecase usecase.User) *UserHandler {
	return &UserHandler{userUsecase}
}
