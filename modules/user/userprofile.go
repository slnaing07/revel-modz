package user

import (
	"time"

	"github.com/jinzhu/gorm"
)

type UserProfileElement struct {
	// gorm fields
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	UserId  int64
	Field   string
	Content string
}

func addAllUserProfileElement(db *gorm.DB, u *UserProfileElement) error {
	// TODO:  add check for existance?
	err := db.Save(u).Error
	return err
}

func addUserProfileElement(db *gorm.DB, uId int64, field, content string) error {
	u := &UserProfileElement{
		UserId:  uId,
		Field:   field,
		Content: content,
	}
	err := db.Save(u).Error
	return err
}

func deleteAllUserProfileElements(db *gorm.DB, uId int64) error {
	return db.Where(&UserProfileElement{UserId: uId}).Delete(UserProfileElement{}).Error
}

func deleteUserProfileElement(db *gorm.DB, uId int64, field string) error {
	return db.Where(&UserProfileElement{UserId: uId}).Delete(UserProfileElement{}).Error
}

func getUserProfileElementsByUserId(db *gorm.DB, userid int64) ([]UserProfileElement, error) {
	var user []UserProfileElement
	err := db.Where(&UserProfileElement{UserId: userid}).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func getUserProfileElementByUserIdField(db *gorm.DB, userid int64, field string) (*UserProfileElement, error) {
	var user UserProfileElement
	err := db.Where(&UserProfileElement{UserId: userid}).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
