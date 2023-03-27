package main

import (
	"github.com/gorilla/mux"
	"github.com/rfauzi44/todov2/modules/v1/activity"
	"github.com/rfauzi44/todov2/modules/v1/todo"
)

func NewApp() (*mux.Router, error) {
	mainRoute := mux.NewRouter()

	db := InitDB()
	activity.NewRoute(mainRoute, db)
	todo.NewRoute(mainRoute, db)

	return mainRoute, nil

}
