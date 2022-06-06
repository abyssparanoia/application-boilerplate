package httperror

import (
	"net/http"

	"github.com/pkg/errors"
	"go.uber.org/zap/zapcore"
)

// CodeToLevel ... convert grpc code to zapcore level
func CodeToLevel(code int) zapcore.Level {
	switch code {
	case http.StatusNotFound, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusServiceUnavailable:
		return zapcore.WarnLevel
	case http.StatusInternalServerError, http.StatusTooManyRequests:
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

// ErrToCode ... convert err to code
func ErrToCode(err error) int {
	switch errors.Cause(err).(type) {
	case ServiceUnavailableError:
		return http.StatusServiceUnavailable
	case UnauthorizedError:
		return http.StatusUnauthorized
	case ForbiddenError:
		return http.StatusForbidden
	case BadRequestError:
		return http.StatusBadRequest
	case NotFoundError:
		return http.StatusNotFound
	default:
	}

	return http.StatusInternalServerError
}
