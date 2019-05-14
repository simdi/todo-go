package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	// "github.com/gorilla/mux"

	todo "github.com/simdi/thesmshub/controllers"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// Get the Port from the environment so we can run on any platform
func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "8080"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}

// Define and configure all routes
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON), // Set content type to json
		middleware.Logger,          // Log api request calls
		middleware.DefaultCompress, // Compress results. Mostly gzipping assets and json.
		middleware.RedirectSlashes, // Redirect slashes to no slash url version
		middleware.Recoverer,       // Recover from panic without crashing server
	)

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/api/todos", todo.Routes())
	})

	return router
}

func main() {
	// Declare a new router
	r := Routes()

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route) // Walk and print out all routes
		return nil
	}

	if err := chi.Walk(r, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error()) // Panic if there is an error
	}

	log.Fatal(http.ListenAndServe(GetPort(), r))
}
