package models

import (
	"github.com/revel/revel"
)

type UserSignup struct {
	Email           string
	Password        string
	PasswordConfirm string
}

func (u *UserSignup) Validate(v *revel.Validation) {
	v.Required(u.Email).
		Message("Email required").
		Key("usersignup.Email")
	v.Email(u.Email).
		Message("Valid email required").
		Key("usersignup.Email")
	v.Required(u.Password).
		Message("Password required").
		Key("usersignup.Password")
	v.MinSize(u.Password, 8).
		Message("Password length must be at least 8").
		Key("usersignup.Password")
	v.Required(u.PasswordConfirm == u.Password).
		Message("The passwords do not match.").
		Key("usersignup.PasswordConfirm")
}

type UserMaillist struct {
	Email string
}

func (u *UserMaillist) Validate(v *revel.Validation) {
	v.Required(u.Email).
		Message("Email required").
		Key("usersignup.Email")
	v.Email(u.Email).
		Message("Valid email required").
		Key("usersignup.Email")
}
