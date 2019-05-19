package models

type Todo struct {
	ID          string `json: "id"`
	Title       string `json: "title"`
	Description string `json: "description`
}

type DB struct {
	Items []Todo
}
