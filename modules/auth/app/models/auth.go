package models

type UserAuthInterface interface {
	AuthId() int64
	AuthSecret() []byte
}

type UserAuth struct {
	UserId            int64  `json:"UserId"`
	HashedPassword    []byte ``
	FailedLoginsCount int32  ``
	LockExpiresAt     int64  ``
}

type UserAuthActivate struct {
	UserId int64 `json:"UserId"`

	ActivatePasswordToken          string ``
	ActivatePasswordTokenExpiresAt int64  ``
	ActivatePasswordEmailSentAt    int64  ``
}

type UserAuthReset struct {
	UserId int64 `json:"UserId"`

	ResetPasswordToken          string ``
	ResetPasswordTokenExpiresAt int64  ``
	ResetPasswordEmailSentAt    int64  ``
}
