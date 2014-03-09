package auth

import (
	"errors"
	"time"

	"code.google.com/p/go.crypto/bcrypt"
	"github.com/jinzhu/gorm"
)

type UserAuth struct {
	// gorm fields
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time

	UserId         int64  `sql:"not null;unique"`
	HashedPassword []byte `sql:"not null"`

	Activated bool
	ResetPass bool

	LastLoginAt  time.Time
	LastLogoutAt time.Time

	FailedLoginsCount int32
	LockExpiresAt     time.Time
}

func checkUserExistsById(db *gorm.DB, uId int64) (bool, error) {
	var ua UserAuth
	err := db.Where(&UserAuth{UserId: uId}).Find(&ua).Error
	if err == gorm.RecordNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	if uId == ua.UserId {
		return true, nil
	}
	return false, nil
}

func addUser(db *gorm.DB, uId int64, password string) error {
	hPass, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	ua := &UserAuth{
		UserId:         uId,
		HashedPassword: hPass,
	}
	found, err := checkUserExistsById(db, uId)
	if err != nil {
		return err
	}

	if !found {
		return db.Save(ua).Error
	} else {
		return errors.New("user already exists")
	}
}
