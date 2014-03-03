package analytics

import (
	"time"
)

type VisiterInfo struct {
	// gorm fields
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time

	VisiterId int64
}

type VisiterPageRequest struct {
	// gorm fields
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time

	VisiterId  int
	Time       time.Time
	Method     string
	RequestURI string
	XRealIp    string
	Referer    string
}

type VisiterMouseEvents struct {
	// gorm fields
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time

	VisiterId int
	Time      time.Time
	Event     string
	Details   string
}

type VisiterKeyboardEvents struct {
	// gorm fields
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time

	VisiterId int
	Time      time.Time
	Event     string
	Details   string
}

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
