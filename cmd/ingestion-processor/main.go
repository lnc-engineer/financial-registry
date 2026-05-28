package main

import (
	"fmt"
	"net/http"

	"github.com/lnc-engineer/financial-registry/internal/middleware"
)

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