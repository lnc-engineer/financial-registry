package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/lnc-engineer/financial-registry/internal/execution"
)

func ExecutionContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		execution.RecordRequest()

		execCtx := execution.ExecutionContext{
			RequestID: uuid.NewString(),
			TraceID:   uuid.NewString(),
			StartTime: time.Now(),
			Metadata:  make(map[string]string),
		}

		ctx := context.WithValue(
			r.Context(),
			execution.ExecutionContextKey,
			execCtx,
		)

		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
