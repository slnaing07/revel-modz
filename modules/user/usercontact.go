package user

import (
	"time"

	"github.com/jinzhu/gorm"
)

type UserAddress struct {
	// gorm fields
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	UserId      int64
	AddressType string // Home, Work, etc

	AddressLine1 string
	AddressLine2 string
	City         string
	State        string
	Zip          string
	Country      string
}

type UserPhone struct {
	// gorm fields
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	UserId      int64
	PhoneType   string // Home, Work, Cell, etc
	PhoneNumber string
}

func addUserAddress(db *gorm.DB, u *UserAddress) error {
	return db.Save(u).Error
}

func addUserPhone(db *gorm.DB, uId int64, p_type, number string) error {
	u := &UserPhone{
		UserId:      uId,
		PhoneType:   p_type,
		PhoneNumber: number,
	}
	return db.Save(u).Error
}

func deleteUserAddress(db *gorm.DB, uId int64) error {
	return db.Where(&UserAddress{UserId: uId}).Delete(UserAddress{}).Error
}

func deleteUserPhone(db *gorm.DB, uId int64) error {
	return db.Where(&UserPhone{UserId: uId}).Delete(UserPhone{}).Error
}

func getUserAddressesByUserId(db *gorm.DB, userid int64) ([]UserAddress, error) {
	var user []UserAddress
	err := db.Where(&UserAddress{UserId: userid}).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func getUserAddressByUserIdType(db *gorm.DB, userid int64, a_type string) (*UserAddress, error) {
	var user UserAddress
	err := db.Where(&UserAddress{UserId: userid}).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func getUserPhonesByUserId(db *gorm.DB, userid int64) ([]UserPhone, error) {
	var user []UserPhone
	err := db.Where(&UserPhone{UserId: userid}).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
func getUserPhoneByUserIdType(db *gorm.DB, userid int64, p_type string) (*UserPhone, error) {
	var user UserPhone
	err := db.Where(&UserPhone{UserId: userid}).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
