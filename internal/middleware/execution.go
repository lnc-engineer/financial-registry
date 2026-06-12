package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/lnc-engineer/financial-registry/internal/execution"
)

func ExecutionContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		execution.RecordRequest()

		// CREATE execution context
		execCtx := execution.ExecutionContext{
			RequestID: uuid.NewString(),
			TraceID:   uuid.NewString(),
			StartTime: start,
			Metadata:  make(map[string]string),
		}

		fmt.Printf(
			"[EXECUTION] request=%s trace=%s\n",
			execCtx.RequestID,
			execCtx.TraceID,
		)

		// INJECT into request context
		ctx := context.WithValue(
			r.Context(),
			execution.ExecutionContextKey,
			execCtx,
		)

		next.ServeHTTP(w, r.WithContext(ctx))

		// Middleware duration (end-to-end request time)
		duration := time.Since(start)

		execution.RecordDuration(
			uint64(duration.Milliseconds()),
		)

		fmt.Println("[METRICS] request duration:", duration)
	})
}
