package analytics

import (
	"time"
)

type VisitorPageRequest struct {
	// gorm fields
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time

	VisitorId  int
	Time       time.Time
	Method     string
	RequestURI string
	XRealIp    string
	Referer    string
}

type VisitorMouseEvents struct {
	// gorm fields
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time

	VisitorId int
	Time      time.Time
	Event     string
	Details   string
}

type VisitorKeyboardEvents struct {
	// gorm fields
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time

	VisitorId int
	Time      time.Time
	Event     string
	Details   string
}
