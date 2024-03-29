package model

import "time"

type User struct {
	Id       int        `json:"id" gorm:"column:id;"`
	Account  string     `json:"account" gorm:"column:account;" validate:"required, min=4"`
	Password string     `json:"password" gorm:"column:password;"`
	Hoten    string     `json:"hoten" gorm:"column:hoten;"`
	Avatar   string     `json:"avatar" gorm:"column:avatar;"`
	Age      string     `json:"age" gorm:"column:age"`
	Score    int        `json:"score" gorm:"column:score"`
	CreateAt *time.Time `json:"create_at" gorm:"column:create_at;"`
	UpdateAt *time.Time `json:"update_at,omitempty" gorm:"column:update_at;"`
}

func (User) TableName() string { return "users" }

type CreateUser struct {
	Id       int    `json:"-" gorm:"column:id;"`
	Account  string `json:"account" gorm:"column:account;" validate:"required, min=4"`
	Password string `json:"password" gorm:"column:password;"`
	Hoten    string `json:"hoten" gorm:"column:hoten;"`
	//Avatar   string `json:"avatar" gorm:"column:avatar"`
}

func (CreateUser) TableName() string { return User{}.TableName() }

// login user
type LoginUser struct {
	Account  string `json:"account" gorm:"column:account;" validate:"required, min=4"`
	Password string `json:"password" gorm:"column:password;"`
}
type ByIdUser struct {
	Id int `json:"id"`
}
type UserResponse struct {
	Id       int        `json:"id"`
	Account  string     `json:"account"`
	Password string     `json:"-" `
	Hoten    string     `json:"hoten"`
	Avatar   string     `json:"avatar,omitempty"`
	Score    int        `json:"score"`
	Age      string     `json:"age"`
	CreateAt *time.Time `json:"create_at"`
	UpdateAt *time.Time `json:"update_at,omitempty"`
}

func (UserResponse) TableName() string { return User{}.TableName() }

type UpdateUser struct {
	//Password string `json:"password" gorm:"column:password;"`
	Hoten    string `json:"hoten" gorm:"column:hoten;"`
	Age      string `json:"age" gorm:"column:age;"`
	Password string `json:"password" gorm:"column:password;"`
	//Avatar string `json:"avatar" gorm:"column:avatar;"`
}

func (UpdateUser) TableName() string { return User{}.TableName() }

type UpdateScoreUser struct {
	Score int `json:"score" gorm:"column:score"`
}

func (UpdateScoreUser) TableName() string { return User{}.TableName() }
