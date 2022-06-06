package server_error

import "github.com/abyssparanoia/application-boilerplate/internal/pkg/error/httperror"

var (
	InternalError   httperror.InternalError     = "InternalErr"
	NotFoundErr     httperror.NotFoundError     = "NotFoundErr"
	UnauthorizedErr httperror.UnauthorizedError = "UnauthorizedErr"
)
