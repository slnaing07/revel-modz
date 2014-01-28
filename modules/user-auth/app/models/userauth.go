package models

type UserAuth struct {
	UserId                      int64  `json:"UserId"`
	HashedPassword              []byte ``
	ResetPasswordToken          string ``
	ResetPasswordTokenExpiresAt int64  ``
	ResetPasswordEmailSentAt    int64  ``
	FailedLoginsCount           int32  ``
	LockExpiresAt               int64  ``
}
