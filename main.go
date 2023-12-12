package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ross96d/simple_swagger_server/handlers"
)

func main() {
	router := chi.NewMux()
	router.Use(middleware.DefaultLogger)

	api := humachi.New(router, huma.DefaultConfig("Chinook api", "1.0.0"))

	handlers.Register(api, "artist")

	fmt.Printf("listen on http://localhost:8888\n")
	if err := http.ListenAndServe(":8888", router); err != nil {
		fmt.Printf("Error running the server %v", err)
		os.Exit(1)
	}
}
