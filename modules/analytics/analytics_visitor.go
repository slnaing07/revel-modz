package analytics

import (
	"time"
)

type VisitorPageRequest struct {
	// gorm fields
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time

	VisitorId  int64
	Time       time.Time
	Method     string
	RequestURI string
	Host       string
	XRealIp    string
	Referer    string
}

type VisitorMouseEvents struct {
	// gorm fields
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time

	VisitorId int64
	Time      time.Time
	Event     string
	Details   string
}

type VisitorKeyboardEvents struct {
	// gorm fields
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time

	VisitorId int64
	Time      time.Time
	Event     string
	Details   string
}
