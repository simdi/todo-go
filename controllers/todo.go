package todo

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type Todo struct {
	Slug  string `json: "slug"`
	Title string `json: "title"`
	Body  string `json: "body`
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", GetAllTodos)
	router.Get("/{todoID}", GetATodo)

	return router
}

func GetATodo(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	todos := Todo{
		Slug:  todoID,
		Title: "Hello World",
		Body:  "Hello from Planet Earth",
	}
	render.JSON(w, r, todos)
}
func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	todos := Todo{
		Slug:  todoID,
		Title: "Hello World",
		Body:  "Hello from Planet Earth",
	}
	render.JSON(w, r, todos)
}
