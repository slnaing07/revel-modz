package models

type UserVolatile struct {
	UserId         int64 `json:"UserId"`
	CreatedAt      int64 `json:"CreatedAt"`
	UpdatedAt      int64 `json:"UpdatedAt"`
	LastLoginAt    int64 `json:"LastLoginAt"`
	LastLogoutAt   int64 `json:"LastLogoutAt"`
	LastActivityAt int64 `json:"LastActivityAt"`
}
