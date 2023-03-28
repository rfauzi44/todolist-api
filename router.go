package main

import (
	"github.com/gorilla/mux"
	"github.com/rfauzi44/todolist-api/module/activity"
	"github.com/rfauzi44/todolist-api/module/todo"
)

func NewApp() (*mux.Router, error) {
	mainRoute := mux.NewRouter()

	db := InitDB()
	activity.NewRoute(mainRoute, db)
	todo.NewRoute(mainRoute, db)

	return mainRoute, nil

}
