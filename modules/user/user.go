package user

import (
	"time"

	// "github.com/jinzhu/gorm"
	// "github.com/robfig/revel"
)

type UserInterface interface {
	GetUserId() int64
	GetUserName() string
}

// Convenience struct for the User/Auth interfaces
// Not stored in DB
type UserPass struct {
	UserId   int64
	UserName string
	Password string
}

func (u UserPass) AuthId() int64 {
	return u.UserId
}

func (u UserPass) AuthPass() string {
	return u.Password
}

func (u UserPass) GetUserId() int64 {
	return u.UserId
}

func (u UserPass) GetUserName() string {
	return u.UserName
}

// Union of all the User* tables
type User struct {
	UserId int64

	Basic   UserBasic
	Detail  UserDetail
	Contact UserContact
	Profile []UserProfileElement
}

type UserVolatile struct {
	Id             int64
	UserId         int64
	LastActivityAt time.Time

	ActivityType string
}

type UserDetail struct {
	Id        int64
	UserId    int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	Title     string
	FirstName string
	Middle    string
	LastName  string

	Dob time.Time
	Sex string
}

type UserContact struct {
	Id        int64
	UserId    int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	ContactType string // Home, Work, etc

	AddressLine1 string
	AddressLine2 string
	City         string
	State        string
	Zip          string
	Country      string
	Phone        string
}

type UserProfileElement struct {
	Id        int64
	UserId    int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	Field   string
	Content string
}
