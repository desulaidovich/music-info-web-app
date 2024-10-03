package middleware

import (
	"app/internal/logger"
	"context"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
)

type requestConextKey string

const requestID requestConextKey = "traceID"

func RequestID(ctx context.Context) string {
	if id, ok := ctx.Value(requestID).(string); ok {
		return id
	}

	return "empty request id"
}

func Logger(next http.Handler, logger *logger.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		ctx := context.WithValue(r.Context(), requestID, id)

		r = r.WithContext(ctx)

		logger.Info(id,
			slog.String("Method", r.Method),
			slog.String("Path", r.RequestURI),
		)

		next.ServeHTTP(w, r)
	})
}
