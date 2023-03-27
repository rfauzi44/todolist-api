package activity

import (
	"fmt"

	"github.com/rfauzi44/todolist-api/database/orm/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type activity_repo struct {
	database *gorm.DB
}

func NewRepo(db *gorm.DB) *activity_repo {
	return &activity_repo{db}

}

func (r *activity_repo) Add(data *model.Activity) (*model.Activity, error) {

	err := r.database.Create(data).Error

	if err != nil {
		return nil, err
	}

	return data, nil

}

func (r *activity_repo) GetAll() (*model.Activities, error) {
	var data model.Activities
	err := r.database.Find(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil

}

func (r *activity_repo) GetById(id int) (*model.Activity, error) {
	var data model.Activity

	err := r.database.First(&data, "activity_id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *activity_repo) Update(id int, data *model.Activity) (*model.Activity, error) {

	err := r.database.Model(&data).Clauses(clause.Returning{}).Where("activity_id = ?", id).Updates(data).Error

	if err != nil {
		return nil, err
	}

	// Fetch the updated instance from the database
	updatedData := &model.Activity{}
	if err := r.database.First(updatedData, id).Error; err != nil {
		return nil, err
	}

	return updatedData, nil
}

func (r *activity_repo) Delete(id int) (map[string]interface{}, error) {

	result := r.database.Delete(&model.Activity{}, id)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("record not found")
	}

	return map[string]interface{}{}, nil
}
