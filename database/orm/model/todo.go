package model

import "time"

type Todo struct {
	ID         int       `gorm:"column:todo_id;primary_key;auto_increment" json:"id,omitempty" valid:"-"`
	ActivityID int       `json:"activity_group_id" valid:"-"`
	Activity   Activity  `gorm:"foreignKey:ActivityID" json:"-" valid:"-"`
	Title      string    `json:"title" valid:"required"`
	IsActive   bool      `gorm:"default:true" json:"is_active" valid:"-"`
	Priority   string    `gorm:"default:very-high" json:"priority,omitempty" valid:"-"`
	CreatedAt  time.Time `json:"createdAt" valid:"-"`
	UpdatedAt  time.Time `json:"updatedAt,omitempty" valid:"-" `
}

type Todos []Todo
