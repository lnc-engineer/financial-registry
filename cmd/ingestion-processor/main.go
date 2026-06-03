package main

import (
	"fmt"
	"github.com/lnc-engineer/financial-registry/internal/execution"
	"github.com/lnc-engineer/financial-registry/internal/middleware"
	"net/http"
	"time"
)

const Version = "0.1.0"

func main() {

	http.Handle(
		"/",
		middleware.LoggingMiddleware(http.HandlerFunc(homeHandler)),
	)

	http.Handle(
		"/health",
		middleware.LoggingMiddleware(http.HandlerFunc(healthHandler)),
	)

	http.Handle(
		"/version",
		middleware.LoggingMiddleware(http.HandlerFunc(versionHandler)),
	)

	http.Handle(
		"/metrics",
		middleware.LoggingMiddleware(http.HandlerFunc(metricsHandler)),
	)

	http.Handle(
		"/process",
		middleware.ExecutionContextMiddleware(
			middleware.LoggingMiddleware(
				http.HandlerFunc(processHandler),
			),
		),
	)

	fmt.Println("Server started on :8080")

	go func() {
		time.Sleep(5 * time.Second)
		execution.PrintMetrics()
	}()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("server error:", err)
	}
}
