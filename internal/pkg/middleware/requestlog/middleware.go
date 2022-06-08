package requestlog

import (
	"fmt"
	"net/http"
	"time"

	"github.com/abyssparanoia/application-boilerplate/internal/pkg/glueerr"
	"github.com/abyssparanoia/application-boilerplate/internal/pkg/renderer"

	"github.com/blendle/zapdriver"
	"github.com/google/uuid"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
)

const producerID = "rapid-go"

// HTTPMiddleware ... http middleware
type HTTPMiddleware struct {
	logger *zap.Logger
}

// Handle ... handle http request
func (m *HTTPMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		operationID := uuid.New()

		ctx := r.Context()
		ctx = ctxzap.ToContext(ctx, m.logger.With(
			zapdriver.OperationCont(operationID.String(), producerID),
		))

		defer func() {
			if rcvr := recover(); rcvr != nil {
				latency := time.Since(start)
				zapcoreLevel := glueerr.CodeToLevel(http.StatusInternalServerError)
				renderer.Text(ctx, w, http.StatusInternalServerError, "Internal Error")
				ctxzap.Extract(ctx).Check(zapcoreLevel, fmt.Sprintf("%s %s -> %d", r.Method, r.RequestURI, http.StatusInternalServerError)).Write(
					zap.Reflect("error", rcvr),
					zap.Int("status", http.StatusInternalServerError),
					zap.Int("content-length", 0),
					zap.Duration("took", latency),
					zap.Int64("latency", latency.Nanoseconds()),
					zap.String("remote", r.RemoteAddr),
					zap.String("request", r.RequestURI),
					zap.String("method", r.Method),
					zap.String("operation-id", operationID.String()),
					zapdriver.OperationEnd(operationID.String(), producerID))
			}
		}()

		sw := &statusWriter{ResponseWriter: w}

		next.ServeHTTP(sw, r.WithContext(ctx))

		latency := time.Since(start)
		zapcoreLevel := glueerr.CodeToLevel(sw.status)
		ctxzap.Extract(ctx).Check(zapcoreLevel, fmt.Sprintf("%s %s -> %d", r.Method, r.RequestURI, sw.status)).Write(
			zap.Int("status", sw.status),
			zap.Int("content-length", sw.length),
			zap.Duration("took", latency),
			zap.Int64("latency", latency.Nanoseconds()),
			zap.String("remote", r.RemoteAddr),
			zap.String("request", r.RequestURI),
			zap.String("method", r.Method),
			zap.String("operation-id", operationID.String()),
			zapdriver.OperationEnd(operationID.String(), producerID))
	})
}

// New ... new http middleware
func New(logger *zap.Logger) *HTTPMiddleware {
	return &HTTPMiddleware{logger}
}
