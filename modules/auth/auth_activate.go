package auth

import (
	"time"

	"github.com/jinzhu/gorm"
)

type UserAuthActivate struct {
	// gorm fields
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time

	UserId                        int64
	ActivateAccountToken          string
	ActivateAccountTokenExpiresAt time.Time
	ActivateAccountEmailSentAt    time.Time
}

func addActivationRecord(db *gorm.DB, uId int64) error {
	return nil
}
