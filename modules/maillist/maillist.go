package maillist

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
)

type MaillistUser struct {
	// gorm fields
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time

	UserId    int64 `sql:"not null;unique"`
	Email     string
	List      string
	Activated bool
}

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

func AddMaillistUser(db *gorm.DB, uId int64, email string) (*MaillistUser, error) {

	mu := &MaillistUser{
		UserId: uId,
		Email:  email,
	}

	if !checkUserExistsById(db, uId) && !checkUserExistsByEmail(db, email) {
		err := db.Save(mu).Error
		if err != nil {
			revel.TRACE.Println("error saving user: ", err)
			return nil, err
		}
	}

	return mu, nil

}

func checkUserExistsById(db *gorm.DB, uId int64) bool {
	var mu MaillistUser
	err := db.Where(&MaillistUser{UserId: uId}).Find(&mu).Error
	if err == gorm.RecordNotFound {
		return false
	}

	if err != nil {
		revel.TRACE.Println("Error looking up user", err)
		return false
	}

	// TODO make this check better
	if uId == mu.UserId {
		return true
	}

	return false
}

func checkUserExistsByEmail(db *gorm.DB, email string) bool {
	var mu MaillistUser
	err := db.Where(&MaillistUser{Email: email}).Find(&mu).Error
	if err == gorm.RecordNotFound {
		return false
	}

	if err != nil {
		revel.TRACE.Println("Error looking up user", err)
		return false
	}

	// TODO make this check better
	if email == mu.Email {
		return true
	}

	return false
}
