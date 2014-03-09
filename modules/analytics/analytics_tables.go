package analytics

import (
	"errors"

	"github.com/jinzhu/gorm"
)

func AddTables(db *gorm.DB) error {
	var err error
	err = db.AutoMigrate(VisitorPageRequest{}).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(VisitorMouseEvents{}).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(VisitorKeyboardEvents{}).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(UserPageRequest{}).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(UserMouseEvents{}).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(UserKeyboardEvents{}).Error
	if err != nil {
		return err
	}
	return nil
}

func DropTables(db *gorm.DB) error {
	var err error
	err = db.DropTable(VisitorPageRequest{}).Error
	if err != nil {
		return err
	}
	err = db.DropTable(VisitorMouseEvents{}).Error
	if err != nil {
		return err
	}
	err = db.DropTable(VisitorKeyboardEvents{}).Error
	if err != nil {
		return err
	}
	err = db.DropTable(UserPageRequest{}).Error
	if err != nil {
		return err
	}
	err = db.DropTable(UserMouseEvents{}).Error
	if err != nil {
		return err
	}
	err = db.DropTable(UserKeyboardEvents{}).Error
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
