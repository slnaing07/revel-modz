package maillist

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/robfig/revel"
)

type MaillistUser struct {
	Id        int64
	UserId    int64 `sql:"not null;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	Email     string
	Activated bool
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
