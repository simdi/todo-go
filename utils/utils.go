package utils

import (
	"github.com/simdi/todo-go/models"
	log "github.com/sirupsen/logrus"
)

// Return a pointer to a new Todo
func New() *models.DB {
	return &models.DB{}
}

func (db *models.DB) Add(item models.Todo) {
	db.Items = append(db.Items, item)
}

// Read data from in memory JSON DB
func ReadFromJSONDB() {
	log.Info("Try to read from JSON DB Store")
}

// Write data to in memory JSON DB
func WriteToJSONDB() {
	log.Info("Try to write to JSON DB Store")
}
