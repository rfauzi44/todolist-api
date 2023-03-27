package todo

import (
	"fmt"

	"github.com/rfauzi44/todolist-api/database/orm/model"
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

	err := r.database.Create(data).Error

	if err != nil {
		return nil, err
	}

	return data, nil

}

func (r *todo_repo) GetAll() (*model.Todos, error) {
	var data model.Todos
	err := r.database.Find(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil

}

func (r *todo_repo) GetById(id int) (*model.Todo, error) {
	var data model.Todo

	err := r.database.First(&data, "todo_id = ?", id).Error
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

	result := r.database.Delete(&model.Todo{}, id)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("record not found")
	}

	return map[string]interface{}{}, nil
}

func (r *todo_repo) Sort(activity_group_id int) (*model.Todos, error) {
	var data model.Todos

	err := r.database.Find(&data, "activity_id = ?", activity_group_id).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}
