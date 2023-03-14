package model

type Filter struct {
	Genre string `json:"genre" form:"genre"`
	Limit int    `json:"limit" form:"limit"`
}
