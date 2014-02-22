package auth

import (
	"time"

	"code.google.com/p/go.crypto/bcrypt"
	"github.com/jinzhu/gorm"
	"github.com/robfig/revel"
)

type UserAuthInterface interface {
	AuthId() int64
	AuthSecret() string
}

type UserAuth struct {
	Id        int64
	UserId    int64 `sql:"not null;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time

	HashedPassword []byte `sql:"not null"`

	Activated bool
	ResetPass bool

	LastLoginAt  time.Time
	LastLogoutAt time.Time

	FailedLoginsCount int32
	LockExpiresAt     time.Time
}

type UserAuthActivate struct {
	Id        int64
	UserId    int64
	CreatedAt time.Time
	UpdatedAt time.Time

	ActivateAccountToken          string
	ActivateAccountTokenExpiresAt time.Time
	ActivateAccountEmailSentAt    time.Time
}

type UserAuthReset struct {
	Id        int64
	UserId    int64
	CreatedAt time.Time
	UpdatedAt time.Time

	ResetPasswordToken          string
	ResetPasswordTokenExpiresAt time.Time
	ResetPasswordEmailSentAt    time.Time
}

func CheckUserAuth(db *gorm.DB, user UserAuthInterface) *UserAuth {
	var ua UserAuth
	err := db.Where(&UserAuth{UserId: user.AuthId()}).Find(&ua).Error
	if err != nil {
		revel.ERROR.Println("Error looking up user", err)
	}

	err = bcrypt.CompareHashAndPassword(ua.HashedPassword, []byte(user.AuthSecret()))
	if err == nil {
		return &ua
	} else {
		revel.ERROR.Println(string(user.AuthSecret()))
		revel.ERROR.Println(string(ua.HashedPassword))
		return nil
	}
}
