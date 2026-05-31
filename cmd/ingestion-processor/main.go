package main

import (
	"fmt"
	"net/http"

	"github.com/lnc-engineer/financial-registry/internal/middleware"
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
		"/process",
		middleware.ExecutionContextMiddleware(
			middleware.LoggingMiddleware(
				http.HandlerFunc(processHandler),
			),
		),
	)

	fmt.Println("Server started on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("server error:", err)
	}
}
