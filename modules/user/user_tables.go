package user

import (
	"errors"

	"github.com/jinzhu/gorm"
)

func AddTables(db *gorm.DB) error {
	var err error
	err = db.AutoMigrate(UserBasic{}).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(UserAddress{}).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(UserPhone{}).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(UserProfileElement{}).Error
	if err != nil {
		return err
	}
	return nil
}

func DropTables(db *gorm.DB) error {
	return db.DropTable(UserBasic{}).Error
	return db.DropTable(UserAddress{}).Error
	return db.DropTable(UserPhone{}).Error
	return db.DropTable(UserProfileElement{}).Error
}

func FillTables(db *gorm.DB) error {
	return errors.New("TODO")
}
func TestTables(db *gorm.DB) error {
	return errors.New("TODO")
}
