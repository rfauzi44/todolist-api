package model

import (
	"time"
)

type Activity struct {
	ID        int       `gorm:"column:activity_id;primary_key;auto_increment" json:"id,omitempty" valid:"-"`
	Title     string    `json:"title" valid:"required"`
	Email     string    `json:"email" valid:"-"`
	CreatedAt time.Time `json:"createdAt" valid:"-"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" valid:"-" `
	TodoList  Todos     `json:"todo_list,omitempty" valid:"-" `
}

type Activities []Activity
