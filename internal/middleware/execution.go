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

		rootSpan := execution.ExecutionContext{
			RequestID:    uuid.NewString(),
			TraceID:      uuid.NewString(),
			SpanID:       uuid.NewString(),
			ParentSpanID: "",
			SpanName: "request",
			StartTime:    time.Now(),
			Metadata:     make(map[string]string),
		}

		execution.LogSpan("ROOT", rootSpan)

		ctx := context.WithValue(
			r.Context(),
			execution.ExecutionContextKey,
			rootSpan,
		)

		// Create CHILD span from root
		childSpan := execution.NewChildSpan(rootSpan)
		childSpan.SpanName = "processing"

		execution.LogSpan("CHILD", childSpan)

		// Override context with child span (downstream execution uses this)
		ctx = context.WithValue(
			ctx,
			execution.ExecutionContextKey,
			childSpan,
		)

		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
