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

		// Create execution context (value, not pointer)
		execCtx := execution.ExecutionContext{
			RequestID: uuid.NewString(),
			StartTime: time.Now(),
		}

		// Inject into request context
		ctx := context.WithValue(
			r.Context(),
			execution.ExecutionContextKey,
			execCtx,
		)

		// Pass request forward with enriched context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}