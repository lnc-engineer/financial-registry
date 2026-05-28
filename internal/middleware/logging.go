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

		fmt.Println("[REQUEST]")
		fmt.Println("method =", r.Method)
		fmt.Println("path =", r.URL.Path)
		fmt.Println("status =", recorder.StatusCode)
		fmt.Println("duration =", duration)
	})
}