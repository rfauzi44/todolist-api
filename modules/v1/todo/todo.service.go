package todo

import (
	"fmt"
	"strings"

	"github.com/rfauzi44/todolist-api/database/orm/model"
	"github.com/rfauzi44/todolist-api/interfaces"
	"github.com/rfauzi44/todolist-api/lib"
)

type todo_service struct {
	repo interfaces.TodoRepoIF
}

func NewService(repo interfaces.TodoRepoIF) *todo_service {
	return &todo_service{repo}

}

func (s *todo_service) Add(data *model.Todo) *lib.Response {
	data, err := s.repo.Add(data)
	if err != nil {
		if strings.Contains(err.Error(), "1452") {
			return lib.NewRes("activity_group_id cannot be null", 400, true)
		}

		return lib.NewRes(err.Error(), 404, true)
	}
	return lib.NewRes(data, 200, false)

}

func (s *todo_service) GetAll() *lib.Response {
	data, err := s.repo.GetAll()
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	return lib.NewRes(data, 200, false)

}

func (s *todo_service) GetById(id int) *lib.Response {
	data, err := s.repo.GetById(id)
	if err != nil {

		if err.Error() == "record not found" {

			return lib.NewRes(fmt.Sprintf("Todo with ID %d Not Found", id), 404, true)

		}

		return lib.NewRes(err.Error(), 400, true)
	}

	return lib.NewRes(data, 200, false)
}

func (s *todo_service) Update(id int, data *model.Todo) *lib.Response {
	data, err := s.repo.Update(id, data)
	if err != nil {

		if err.Error() == "record not found" {

			return lib.NewRes(fmt.Sprintf("Todo with ID %d Not Found", id), 404, true)

		}

		return lib.NewRes(err.Error(), 400, true)
	}
	return lib.NewRes(data, 200, false)

}

func (s *todo_service) Delete(id int) *lib.Response {
	data, err := s.repo.Delete(id)

	if err != nil {

		if err.Error() == "record not found" {

			return lib.NewRes(fmt.Sprintf("Todo with ID %d Not Found", id), 404, true)

		}

		return lib.NewRes(err.Error(), 400, true)
	}
	return lib.NewRes(data, 200, false)

}

func (s *todo_service) Sort(activity_group_id int) *lib.Response {
	data, err := s.repo.Sort(activity_group_id)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}

	return lib.NewRes(data, 200, false)
}
