package models

import (
	"time"
)

type User struct {
	UserBase
}

type UserBase struct {
	UserId   int64
	UserName string

	Volt    UserVolatile
	Detail  UserDetail
	Contact UserContact
	Profile UserProfile
	Social  UserSocial
}

type UserVolatile struct {
	UserId         int64
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
	LastLoginAt    time.Time
	LastLogoutAt   time.Time
	LastActivityAt time.Time
}

type UserDetail struct {
	UserId int64

	Title         string
	FirstName     string
	LastName      string
	MiddleInitial string

	Dob string
	Sex string
}

type UserContact struct {
	UserId int64

	ContactType  string // Home, Work, etc
	AddressLine1 string
	AddressLine2 string
	City         string
	State        string
	Zip          string
	Email        string
	Phone        string
}

type UserProfile struct {
}

type UserSocial struct {
}

func (u UserBase) AuthId() string {
	return u.Identifier
}
func (u UserBase) AuthSecret() string {
	return string(u.HashPass)
}
