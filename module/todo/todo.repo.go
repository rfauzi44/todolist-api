package todo

import (
	"fmt"

	"github.com/rfauzi44/todolist-api/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type todo_repo struct {
	database *gorm.DB
}

func NewRepo(db *gorm.DB) *todo_repo {
	return &todo_repo{db}

}

func (r *todo_repo) Add(data *model.Todo) (*model.Todo, error) {

	session := r.database.Session(&gorm.Session{PrepareStmt: true})
	tx := session.Begin()

	err := tx.Create(data).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	return data, nil

}

func (r *todo_repo) GetAll() (*model.Todos, error) {
	var data model.Todos
	session := r.database.Session(&gorm.Session{PrepareStmt: true})
	err := session.Find(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil

}

func (r *todo_repo) GetById(id int) (*model.Todo, error) {
	var data model.Todo
	session := r.database.Session(&gorm.Session{PrepareStmt: true})

	err := session.First(&data, "todo_id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *todo_repo) Update(id int, data *model.Todo) (*model.Todo, error) {
	err := r.database.Model(&data).Clauses(clause.Returning{}).Where("todo_id = ?", id).Updates(data).Error

	if err != nil {
		return nil, err
	}

	// Fetch the updated instance from the database
	updatedData := &model.Todo{}
	if err := r.database.First(updatedData, id).Error; err != nil {
		return nil, err
	}

	return updatedData, nil
}

func (r *todo_repo) Delete(id int) (map[string]interface{}, error) {

	session := r.database.Session(&gorm.Session{PrepareStmt: true})
	tx := session.Begin()

	result := tx.Delete(&model.Todo{}, id)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		tx.Rollback()
		return nil, fmt.Errorf("record not found")
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return map[string]interface{}{}, nil
}

func (r *todo_repo) Sort(activity_group_id int) (*model.Todos, error) {
	var data model.Todos
	session := r.database.Session(&gorm.Session{PrepareStmt: true})

	err := session.Find(&data, "activity_group_id = ?", activity_group_id).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}
