package middleware

import (
	"fmt"
	"net/http"
	"time"
)

type StatusRecorder struct {
	http.ResponseWriter
	StatusCode int
}

func (sr *StatusRecorder) WriteHeader(statusCode int) {
	sr.StatusCode = statusCode
	sr.ResponseWriter.WriteHeader(statusCode)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		recorder := &StatusRecorder{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}

		next.ServeHTTP(recorder, r)

		duration := time.Since(start)

		fmt.Printf(
			`{"type":"request","method":"%s","path":"%s","status":%d,"duration_us":%d}\n`,
			r.Method,
			r.URL.Path,
			recorder.StatusCode,
			duration.Microseconds(),
		)
	})
}