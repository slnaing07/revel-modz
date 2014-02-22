package controllers

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/robfig/revel"
)

type UserBasic struct {
	Id     int64 // Primary Key form Gorm
	UserId int64 // Unique identifier per user across all tables
	Email  string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetUserById(db *gorm.DB, id int) *UserBasic {
	var user UserBasic
	err := db.Where(&UserBasic{Id: int64(id)}).First(&user).Error
	// TODO: change this to check error type
	if err != nil {
		revel.TRACE.Println(err)
		return nil
	}

	// user.Id ~= 0 if no record found, related to TODO
	if user.Id != int64(id) {
		return nil
	}

	return &user
}

func GetUserByUserId(db *gorm.DB, userid int) *UserBasic {
	var user UserBasic
	err := db.Where(&UserBasic{UserId: int64(userid)}).First(&user).Error
	// TODO: change this to check error type
	if err != nil {
		revel.TRACE.Println(err)
		return nil
	}

	// user.UserId ~= 0 if no record found, related to TODO
	if user.UserId != int64(userid) {
		return nil
	}

	return &user
}

func GetUserByEmail(db *gorm.DB, email string) *UserBasic {
	var user UserBasic
	err := db.Where(&UserBasic{Email: email}).First(&user).Error
	// TODO: change this to check error type
	if err != nil {
		revel.TRACE.Println(err)
		return nil
	}

	if user.Email != email {
		return nil
	}

	return &user
}
