package activity

import (
	"fmt"

	"github.com/rfauzi44/todolist-api/interfaces"
	"github.com/rfauzi44/todolist-api/lib"
	"github.com/rfauzi44/todolist-api/model"
)

type activity_service struct {
	repo interfaces.ActivityRepoIF
}

func NewService(repo interfaces.ActivityRepoIF) *activity_service {
	return &activity_service{repo}

}

func (s *activity_service) Add(data *model.Activity) *lib.Response {
	data, err := s.repo.Add(data)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	return lib.NewRes(data, 200, false)

}

func (s *activity_service) GetAll() *lib.Response {
	data, err := s.repo.GetAll()
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	return lib.NewRes(data, 200, false)

}

func (s *activity_service) GetById(id int) *lib.Response {
	data, err := s.repo.GetById(id)
	if err != nil {

		if err.Error() == "record not found" {

			return lib.NewRes(fmt.Sprintf("Activity with ID %d Not Found", id), 404, true)

		}

		return lib.NewRes(err.Error(), 400, true)
	}

	return lib.NewRes(data, 200, false)
}

func (s *activity_service) Update(id int, data *model.Activity) *lib.Response {
	data, err := s.repo.Update(id, data)
	if err != nil {

		if err.Error() == "record not found" {

			return lib.NewRes(fmt.Sprintf("Activity with ID %d Not Found", id), 404, true)

		}

		return lib.NewRes(err.Error(), 400, true)
	}
	return lib.NewRes(data, 200, false)

}

func (s *activity_service) Delete(id int) *lib.Response {
	data, err := s.repo.Delete(id)

	if err != nil {

		if err.Error() == "record not found" {

			return lib.NewRes(fmt.Sprintf("Activity with ID %d Not Found", id), 404, true)

		}

		return lib.NewRes(err.Error(), 400, true)
	}
	return lib.NewRes(data, 200, false)

}
