package interfaces

import (
	"github.com/rfauzi44/todolist-api/database/orm/model"
	"github.com/rfauzi44/todolist-api/lib"
)

type ActivityRepoIF interface {
	Add(data *model.Activity) (*model.Activity, error)
	GetAll() (*model.Activities, error)
	GetById(id int) (*model.Activity, error)
	Update(id int, data *model.Activity) (*model.Activity, error)
	Delete(id int) (map[string]interface{}, error)
}

type ActivityServiceIF interface {
	Add(data *model.Activity) *lib.Response
	GetAll() *lib.Response
	GetById(id int) *lib.Response
	Update(id int, data *model.Activity) *lib.Response
	Delete(id int) *lib.Response
}
