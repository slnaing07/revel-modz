package rbac

type SessionInterface interface {
}

type Session struct {
	// Primary Key
	SessionId int64
	UserId    int64

	SessionName string
	CreatedAt   int64 // unix nano
	ExpiresAt   int64 // unix nanoq
}
