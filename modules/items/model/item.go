package model

import "time"

type ToDoItem struct {
	Id       int        `json:"id" gorm:"column:id;"`
	Title    string     `json:"title" gorm:"column:title;" validate:"required, min=4"`
	CreateAt *time.Time `json:"create_at" gorm:"column:create_at;"`
	UpdateAt *time.Time `json:"update_at,omitempty" gorm:"column:update_at;"`
	Image    string     `json:"image" gorm:"column:image;"`
}

func (ToDoItem) TableName() string { return "todo_items" }

type ToDoItemCreate struct {
	Id    int    `json:"-" gorm:"column:id;"`
	Title string `json:"title" gorm:"column:title;" validate:"required, min=4"`
}

func (ToDoItemCreate) TableName() string { return ToDoItem{}.TableName() }

type ToDoItemUpdate struct {
	Title *string `json:"title" gorm:"column:title;"`
}

func (ToDoItemUpdate) TableName() string { return ToDoItem{}.TableName() }

type UploadImage struct {
	Image string `json:"image" gorm:"column:image;"`
}

func (UploadImage) TableName() string { return ToDoItem{}.TableName() }
