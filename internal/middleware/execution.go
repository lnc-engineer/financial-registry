package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/lnc-engineer/financial-registry/internal/execution"
)

var spanBuffer []execution.ExecutionContext

func ExecutionContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		execution.RecordRequest()

		rootSpan := execution.ExecutionContext{
			RequestID:    uuid.NewString(),
			TraceID:      uuid.NewString(),
			SpanID:       uuid.NewString(),
			ParentSpanID: "",
			SpanName:     "request",
			StartTime:    time.Now(),
			Attributes:   map[string]string{},
			Lifecycle:    []execution.LifecycleEvent{},
		}

		rootSpan.AddLifecycleEvent("Span Started")

		// ADD ROOT TO BUFFER
		spanBuffer = append(spanBuffer, rootSpan)

		execution.LogSpan("ROOT", rootSpan)

		ctx := context.WithValue(
			r.Context(),
			execution.ExecutionContextKey,
			rootSpan,
		)

		// Create CHILD span
		childSpan := execution.NewChildSpan(rootSpan, "processing")

		// ADD CHILD TO BUFFER
		spanBuffer = append(spanBuffer, childSpan)

		execution.LogSpan("CHILD", childSpan)

		ctx = context.WithValue(
			ctx,
			execution.ExecutionContextKey,
			childSpan,
		)

		// run actual handler
		next.ServeHTTP(w, r.WithContext(ctx))

		// PRINT TREE AFTER REQUEST FINISHES
		roots := execution.BuildTraceTree(spanBuffer)

		execution.PrintTraceTree(spanBuffer)
		execution.PrintTraceJSON(roots)
		spanBuffer = nil
	})
}
