package todo

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"github.com/rfauzi44/todolist-api/database/orm/model"
	"github.com/rfauzi44/todolist-api/interfaces"
	"github.com/rfauzi44/todolist-api/lib"
)

type todo_controller struct {
	service interfaces.TodoServiceIF
}

func NewController(service interfaces.TodoServiceIF) *todo_controller {
	return &todo_controller{service}

}

func (c *todo_controller) Add(w http.ResponseWriter, r *http.Request) {

	var data model.Todo

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		lib.NewRes(err.Error(), 500, true).Send(w)
		return
	}

	_, err = govalidator.ValidateStruct(data)
	if err != nil {

		if err.Error() == "title: non zero value required" {

			lib.NewRes("title cannot be null", 400, true).Send(w)
			return

		}

		lib.NewRes(err.Error(), 400, true).Send(w)
		return
	}

	result := c.service.Add(&data)

	if result.Code == 200 {

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusCreated)

		result.Send(w)

	} else {

		result.Send(w)

	}

}

func (c *todo_controller) GetAll(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()

	if len(query) == 0 {
		c.service.GetAll().Send(w)

	} else {

		activity_group_id, err := strconv.Atoi(query["activity_group_id"][0])
		if err != nil {
			http.Error(w, "Invalid parameter", http.StatusBadRequest)
			return
		}
		c.service.Sort(activity_group_id).Send(w)

	}

}

func (c *todo_controller) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid parameter", http.StatusBadRequest)
		return
	}

	c.service.GetById(id).Send(w)
}

func (c *todo_controller) Update(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid parameter", http.StatusBadRequest)
		return
	}

	var data model.Todo

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		lib.NewRes(err.Error(), 500, true).Send(w)
		return
	}

	if err != nil {

		if err.Error() == "title: non zero value required" {

			lib.NewRes("title cannot be null ", 400, true).Send(w)
			return

		}

		lib.NewRes(err.Error(), 400, true).Send(w)
		return
	}

	c.service.Update(id, &data).Send(w)

}

func (c *todo_controller) Delete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid parameter", http.StatusBadRequest)
		return
	}

	c.service.Delete(id).Send(w)

}
