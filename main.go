package main

import (
	"fmt"
	"net/http"

	"github.com/pentest-duck/sc-interns-extension/handlers"
)

func main() {
	const port = "8080"

	// Set up HTTP multiplexer
	mux := http.NewServeMux()

	fmt.Println("Server started on port", port)

	// Serve CSS as static file
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Set up handler for idnex
	mux.HandleFunc("/", handlers.IndexHandler)

	// Start web server
	http.ListenAndServe(":"+port, mux)
}
