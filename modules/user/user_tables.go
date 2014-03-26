package user

import (
	"errors"

	"github.com/jinzhu/gorm"
)

func AddTables(db *gorm.DB) error {
	var err error
	err = db.AutoMigrate(Visitor{}).Error
	if err != nil {
		return err
	}
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
	var err error
	err = db.DropTable(Visitor{}).Error
	if err != nil {
		return err
	}
	err = db.DropTable(UserBasic{}).Error
	if err != nil {
		return err
	}
	err = db.DropTable(UserAddress{}).Error
	if err != nil {
		return err
	}
	err = db.DropTable(UserPhone{}).Error
	if err != nil {
		return err
	}
	err = db.DropTable(UserProfileElement{}).Error
	if err != nil {
		return err
	}
	return nil
}

func FillTables(db *gorm.DB) error {
	return errors.New("TODO")
}
func TestTables(db *gorm.DB) error {
	return errors.New("TODO")
}

/**
var fillTables Users = []User{
	User{
		UserId: 100001,
		Basic: UserBasic{

		},
		Detail: UserDetail{

		},
		Addresses: []UserAddress{
			UserAddress{

			}

		}


	}
}

*/
