package maillist

import (
	"github.com/jinzhu/gorm"
)

func AddUser(db *gorm.DB, uId int64, email string) (*MaillistUser, error) {
	mu := &MaillistUser{
		UserId: uId,
		Email:  email,
	}
	err := addUser(db, mu)
	return mu, err
}

func RemoveUser(db *gorm.DB, uId int64) error {
	return removeUserById(db, uId)
}

func GetUserById(db *gorm.DB, uId int64) (*MaillistUser, error) {
	return getUserById(db, uId)
}

func GetUserByEmail(db *gorm.DB, email string) (*MaillistUser, error) {
	return getUserByEmail(db, email)
}

func GetAllUsers(db *gorm.DB) ([]*MaillistUser, error) {
	return getAllUsers(db)
}

func GetUsersByList(db *gorm.DB, list string) ([]*MaillistUser, error) {
	return getUsersByList(db, list)
}
