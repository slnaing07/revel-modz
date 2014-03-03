package user

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
)

type UserBasic struct {
	// gorm fields
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time

	UserId   int64  `sql:"not null;unique"`
	UserName string `sql:"not null;unique"`
}

func AddUserBasic(db *gorm.DB, u *UserBasic) error {
	err := db.Save(u).Error
	return err
}

func GetUserBasicById(db *gorm.DB, id int64) *UserBasic {
	var user UserBasic
	err := db.Where(&UserBasic{Id: id}).First(&user).Error
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

func GetUserBasicByUserId(db *gorm.DB, userid int64) *UserBasic {
	var user UserBasic
	err := db.Where(&UserBasic{UserId: userid}).First(&user).Error
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

func GetUserBasicByName(db *gorm.DB, name string) *UserBasic {
	var user UserBasic
	err := db.Where(&UserBasic{UserName: name}).First(&user).Error
	// TODO: change this to check error type
	if err != nil {
		revel.TRACE.Println(err)
		return nil
	}

	if user.UserName != name {
		return nil
	}

	return &user
}
