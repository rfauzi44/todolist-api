package activity

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewRoute(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/activity-groups").Subrouter()

	repo := NewRepo(db)
	service := NewService(repo)
	controller := NewController(service)

	route.HandleFunc("", controller.Add).Methods("POST")
	route.HandleFunc("", controller.GetAll).Methods("GET")
	route.HandleFunc("/{id}", controller.GetById).Methods("GET")
	route.HandleFunc("/{id}", controller.Update).Methods("PATCH")
	route.HandleFunc("/{id}", controller.Delete).Methods("DELETE")
}
