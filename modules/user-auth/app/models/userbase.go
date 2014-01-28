package models

type UserBase struct {
	UserId   int64  `json:"UserId"`
	UserName string `json:"UserName"`
	Email    string `json:"Email"`
}
