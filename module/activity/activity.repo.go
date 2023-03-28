package activity

import (
	"fmt"

	"github.com/rfauzi44/todolist-api/model"
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

func (r *activity_repo) GetAll() (*model.Activities, error) {
	var data model.Activities
	session := r.database.Session(&gorm.Session{PrepareStmt: true})
	err := session.Find(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil

}

func (r *activity_repo) GetById(id int) (*model.Activity, error) {
	var data model.Activity
	session := r.database.Session(&gorm.Session{PrepareStmt: true})

	err := session.First(&data, "activity_id = ?", id).Error
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
	session := r.database.Session(&gorm.Session{PrepareStmt: true})
	tx := session.Begin()

	result := tx.Delete(&model.Activity{}, id)
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
