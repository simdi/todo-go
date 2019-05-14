package todo

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/simdi/todo-go/utils"
)

type Todo struct {
	ID          string `json: "id"`
	Title       string `json: "title"`
	Description string `json: "description`
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", GetAllTodos)
	router.Post("/", CreateATodo)
	router.Get("/{todoID}", GetATodo)

	return router
}

// Get a single todo from the DB
func GetATodo(w http.ResponseWriter, r *http.Request) {
	utils.ReadFromJSONDB()
	todoID := chi.URLParam(r, "todoID")
	todos := Todo{
		ID:          todoID,
		Title:       "Hello World",
		Description: "Hello from Planet Earth",
	}

	render.JSON(w, r, todos)
}

// Create todo
func CreateATodo(w http.ResponseWriter, r *http.Request) {
	// log.Println(chi)
	// Validate the data
	// Save the data into the database
	// Return an object of the saved data in a json format

	// render.JSON(w, r, todo)
}

// Get all todos from the DB
func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	todos := Todo{
		ID:          todoID,
		Title:       "Hello World",
		Description: "Hello from Planet Earth",
	}
	render.JSON(w, r, todos)
}
