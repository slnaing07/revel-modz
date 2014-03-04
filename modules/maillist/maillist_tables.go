package maillist

import (
	"errors"

	"github.com/jinzhu/gorm"
)

func AddTables(db *gorm.DB) error {
	return db.AutoMigrate(MaillistUser{}).Error
}

func DropTables(db *gorm.DB) error {
	return db.DropTable(MaillistUser{}).Error
}

func FillTables(db *gorm.DB) error {
	return errors.New("TODO")
}
func TestTables(db *gorm.DB) error {
	return errors.New("TODO")
}
