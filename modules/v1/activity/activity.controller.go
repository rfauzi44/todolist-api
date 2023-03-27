package activity

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

type activity_controller struct {
	service interfaces.ActivityServiceIF
}

func NewController(service interfaces.ActivityServiceIF) *activity_controller {
	return &activity_controller{service}

}

func (c *activity_controller) Add(w http.ResponseWriter, r *http.Request) {

	var data model.Activity

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

func (c *activity_controller) GetAll(w http.ResponseWriter, r *http.Request) {
	c.service.GetAll().Send(w)
}

func (c *activity_controller) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		lib.NewRes(err.Error(), 500, true).Send(w)
		return
	}

	c.service.GetById(id).Send(w)
}

func (c *activity_controller) Update(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		lib.NewRes(err.Error(), 500, true).Send(w)
		return
	}

	var data model.Activity

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		lib.NewRes(err.Error(), 500, true).Send(w)
		return
	}

	_, err = govalidator.ValidateStruct(data)
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

func (c *activity_controller) Delete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid parameter", http.StatusBadRequest)
		return
	}

	c.service.Delete(id).Send(w)

}
