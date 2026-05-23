package main

import (
	"fmt"
	"github.com/lnc-engineer/financial-registry/internal/middleware"
	"net/http"
)

func main() {

	http.HandleFunc("/", loggingMiddleware(homeHandler))

	http.Handle("/process", middleware.LoggingMiddleware(http.HandlerFunc(processHandler)))

	http.HandleFunc("/health", loggingMiddleware(healthHandler))

	fmt.Println("Server started on :8080")

	http.ListenAndServe(":8080", nil)
}
