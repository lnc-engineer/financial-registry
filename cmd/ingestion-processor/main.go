package main

import (
	"fmt"
	"net/http"
)


func main() {

	http.HandleFunc("/process", loggingMiddleware(processHandler))
	
	http.HandleFunc("/health", loggingMiddleware(healthHandler))

	fmt.Println("Server started on :8080")

	http.ListenAndServe(":8080", nil)
}
