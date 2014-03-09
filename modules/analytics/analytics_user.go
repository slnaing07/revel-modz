package analytics

import (
	"time"
)

type UserPageRequest struct {
	// gorm fields
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time

	UserId     int
	Time       time.Time
	Method     string
	RequestURI string
	XRealIp    string
	Referer    string
}

type UserMouseEvents struct {
	// gorm fields
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time

	UserId  int
	Time    time.Time
	Event   string
	Details string
}

type UserKeyboardEvents struct {
	// gorm fields
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time

	UserId  int
	Time    time.Time
	Event   string
	Details string
}
