package auth

import (
	"errors"
	"time"

	"code.google.com/p/go.crypto/bcrypt"
	"github.com/jinzhu/gorm"
	"github.com/robfig/revel"
)

type UserAuthInterface interface {
	AuthId() int64
	AuthPass() string
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

func AddUserAuth(db *gorm.DB, user UserAuthInterface) (*UserAuth, error) {

	hPass, _ := bcrypt.GenerateFromPassword([]byte(user.AuthPass()), bcrypt.DefaultCost)

	ua := UserAuth{
		UserId:         user.AuthId(),
		HashedPassword: hPass,
	}

	if !checkUserExistsById(db, user) {
		err := db.Save(ua).Error
		if err != nil {
			return nil, err
		}
	}

	created_at := ua.CreatedAt
	updated_at := ua.UpdatedAt

	if created_at.IsZero() {
		revel.ERROR.Println("Should have created_at after auth create")
	}
	if updated_at.IsZero() {
		revel.ERROR.Println("Should have updated_at after auth create")
	}

	return &ua, nil

}

func Authenticate(db *gorm.DB, user UserAuthInterface) (*UserAuth, error) {
	var ua UserAuth
	err := db.Where(&UserAuth{UserId: user.AuthId()}).Find(&ua).Error
	// TODO: change this to check error type  No Record Found can be returned
	if err != nil {
		revel.ERROR.Println("Error looking up user", err)
		return nil, err
	}

	// TODO make this check better
	// ua.UserId should be 0 when no record found
	if user.AuthId() != int64(ua.UserId) {
		return nil, errors.New("Record Not Found")
	}

	err = bcrypt.CompareHashAndPassword(ua.HashedPassword, []byte(user.AuthPass()))
	if err != nil {
		return nil, errors.New("Password Fail")
	}

	return &ua, nil
}

func checkUserExistsById(db *gorm.DB, user UserAuthInterface) bool {
	var ua UserAuth
	err := db.Where(&UserAuth{UserId: user.AuthId()}).Find(&ua).Error
	if err == gorm.RecordNotFound {
		return false
	}

	if err != nil {
		revel.ERROR.Println("Error looking up user", err)
		return false
	}

	// TODO make this check better
	if user.AuthId() == int64(ua.UserId) {
		return true
	}

	return false
}
