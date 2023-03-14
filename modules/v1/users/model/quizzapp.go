package model

import "time"

type Quiz struct {
	Id       int        `json:"id" gorm:"column:id"`
	Question string     `json:"question" gorm:"column:cau_hoi"`
	Op1      string     `json:"op1" gorm:"column:op1"`
	Op2      string     `json:"op2" gorm:"column:op2"`
	Op3      string     `json:"op3" gorm:"column:op3"`
	Op4      string     `json:"op4" gorm:"column:op4"`
	Ans1     bool       `json:"ans1" gorm:"column:ans1"`
	Ans2     bool       `json:"ans2" gorm:"column:ans2"`
	Ans3     bool       `json:"ans3" gorm:"column:ans3"`
	Ans4     bool       `json:"ans4" gorm:"column:ans4"`
	Genre    string     `json:"mon" gorm:"column:mon"`
	CreateAt *time.Time `json:"create_at" gorm:"column:create_at;"`
	UpdateAt *time.Time `json:"update_at,omitempty" gorm:"column:update_at;"`
}

func (Quiz) TableName() string { return "quizz" }
