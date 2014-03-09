package user

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Visitor struct {
	// gorm fields
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	UserId    int64  // if we know it later
	VisitorId int64  `sql:"not null;unique"`
	VisitorIp string `sql:"not null"`
}

func addVisitor(db *gorm.DB, u *Visitor) error {
	// TODO:  add check for existance?
	err := db.Save(u).Error
	return err
}

func deleteVisitor(db *gorm.DB, uId int64) error {
	return db.Where(&Visitor{UserId: uId}).Delete(Visitor{}).Error
}

func getVisitorById(db *gorm.DB, id int64) (*Visitor, error) {
	var user Visitor
	err := db.Where(&Visitor{Id: id}).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func getVisitorByUserId(db *gorm.DB, id int64) (*Visitor, error) {
	var user Visitor
	err := db.Where(&Visitor{UserId: id}).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func getVisitorByVisitorId(db *gorm.DB, userid int64) (*Visitor, error) {
	var user Visitor
	err := db.Where(&Visitor{VisitorId: userid}).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func getVisitorByVisitorIp(db *gorm.DB, ip string) (*Visitor, error) {
	var user Visitor
	err := db.Where(&Visitor{VisitorIp: ip}).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
