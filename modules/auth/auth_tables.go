package auth

import (
	"errors"

	"github.com/jinzhu/gorm"
)

func AddTables(db *gorm.DB) error {
	return db.AutoMigrate(UserAuth{}).Error
}

func DropTables(db *gorm.DB) error {
	return db.DropTable(UserAuth{}).Error
}

func FillTables(db *gorm.DB) error {
	return errors.New("TODO")
}
func TestTables(db *gorm.DB) error {
	return errors.New("TODO")
}
