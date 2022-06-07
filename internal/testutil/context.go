package testutil

import (
	"context"

	"github.com/abyssparanoia/application-boilerplate/internal/pkg/logger"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
)

var (
	testCtx context.Context
)

func init() {
	logger := logger.New()
	testCtx = ctxzap.ToContext(context.Background(), logger)
}

// Context ... context for test
func Context() context.Context {
	return testCtx
}
