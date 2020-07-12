package grpcerror

import (
	"errors"

	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/codes"
)

// CodeToLevel ... convert grpc code to zapcore level
func CodeToLevel(code codes.Code) zapcore.Level {
	switch code {
	case codes.NotFound, codes.InvalidArgument, codes.AlreadyExists, codes.Unauthenticated, codes.PermissionDenied:
		return zapcore.WarnLevel
	case codes.Internal, codes.Unknown, codes.Aborted:
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

// ErrToCode ... convert err to code
func ErrToCode(err error) codes.Code {
	var notFoundError *NotFoundError
	if ok := errors.As(err, &notFoundError); ok {
		return codes.NotFound
	}
	var invalidArgumentError *InvalidArgumentError
	if ok := errors.As(err, &invalidArgumentError); ok {
		return codes.InvalidArgument
	}
	var internalError *InternalError
	if ok := errors.As(err, &internalError); ok {
		return codes.Internal
	}
	return codes.Unknown
}
