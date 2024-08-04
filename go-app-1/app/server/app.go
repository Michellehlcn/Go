package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewApp() *echo.Echo {
	engine := echo.New()
	engine.Debug = true
	engine.Use(middleware.Recover())

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
