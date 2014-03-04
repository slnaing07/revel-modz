package maillist

import (
	"time"
)

type MaillistMessage struct {
	// gorm fields
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time

	UserId int64 `sql:"not null;unique"`
	Author string

	Draft bool
	Sent  bool

	List      string
	Subject   string
	PlainBody string
	HtmlBody  string
}
