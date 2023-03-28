package interfaces

import (
	"github.com/rfauzi44/todolist-api/lib"
	"github.com/rfauzi44/todolist-api/model"
)

type TodoRepoIF interface {
	Add(data *model.Todo) (*model.Todo, error)
	GetAll() (*model.Todos, error)
	GetById(id int) (*model.Todo, error)
	Update(id int, data *model.Todo) (*model.Todo, error)
	Delete(id int) (map[string]interface{}, error)
	Sort(activity_group_id int) (*model.Todos, error)
}

type TodoServiceIF interface {
	Add(data *model.Todo) *lib.Response
	GetAll() *lib.Response
	GetById(id int) *lib.Response
	Update(id int, data *model.Todo) *lib.Response
	Delete(id int) *lib.Response
	Sort(activity_group_id int) *lib.Response
}
