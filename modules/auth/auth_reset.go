package auth

import (
	"time"

	"github.com/jinzhu/gorm"
)

type UserAuthReset struct {
	// gorm fields
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time

	UserId                      int64
	ResetPasswordToken          string
	ResetPasswordTokenExpiresAt time.Time
	ResetPasswordEmailSentAt    time.Time
}

func addPasswordResetRecord(db *gorm.DB, uId int64) error {
	return nil
}
