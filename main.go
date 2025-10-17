package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Простой HTTP сервер
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Simple Go App</h1><p>CI/CD is working! 🚀</p>")
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	fmt.Println("Server starting on :8080...")
	http.ListenAndServe(":8080", nil)
}
