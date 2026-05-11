package main

import (
	"fmt"
	"net/http"
)


func main() {

	http.HandleFunc("/process", processHandler)

	fmt.Println("Server started on :8080")

	http.ListenAndServe(":8080", nil)
}
