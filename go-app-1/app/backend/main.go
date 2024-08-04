package main

import (
	"net/http"
)

func main() {
	// static files
	fs := http.FileServer(http.Dir("build"))
	http.Handle("/", fs)

	// API routes
	api := http.NewServeMux()
	api.HandleFunc("/api/data", dataHandler)
	http.Handle("/api/", api)

	http.ListenAndServe(":8080", nil)
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	// Handle the request and send the response
	w.Write([]byte("Hello, World!"))

}
