package models

import (
	"regexp"

	"github.com/revel/revel"
)

type UserMaillist struct {
	Email string
}

func (u *UserMaillist) Validate(v *revel.Validation) {
	v.Required(u.Email).
		Message("Email required").
		Key("usermaillist.Email")
	v.Email(u.Email).
		Message("Valid email required").
		Key("usermaillist.Email")
}

type UserLogin struct {
	Email    string
	Password string
}

func (u *UserLogin) Validate(v *revel.Validation) {
	v.Required(u.Email).
		Message("Email required").
		Key("userlogin.Email")
	v.Email(u.Email).
		Message("Valid email address required").
		Key("userlogin.Email")
	v.Required(u.Password).
		Message("Password required").
		Key("userlogin.Password")
}

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

type UserRegister struct {
	Username        string
	Fname           string
	MidInit         string
	Lname           string
	Dob             string
	Sex             string
	Address         string
	City            string
	Zipcode         string
	Phnumber        string
	Email           string
	Password        string
	PasswordConfirm string
}

func (u *UserRegister) Validate(v *revel.Validation) {
	v.Required(u.Username).
		Message("User name required").
		Key("userregister.Username")
	v.MinSize(u.Username, 6).
		Message("Username length must be at least 8").
		Key("userregister.Username")
	v.Required(u.Fname).
		Message("First name required").
		Key("userregister.Fname")
	v.Required(u.Lname).
		Message("Last name required").
		Key("userregister.Lname")
	v.Required(u.Dob).
		Message("Date of Birth required").
		Key("userregister.Dob")
	v.Match(u.Dob, regexp.MustCompile(`(0[1-9]|1[012])[- \/.](0[1-9]|[12][0-9]|3[01])[- \/.](19|20)\d\d`)).
		Message("Date of Birth must be in form MM-DD-YYYY.").
		Key("userregister.Dob")
	v.Required(u.Sex).
		Message("Sex required").
		Key("userregister.Sex")
	v.Required(u.Address).
		Message("Address required").
		Key("userregister.Address")
	v.Required(u.City).
		Message("City required").
		Key("userregister.City")
	v.Match(u.City, regexp.MustCompile(`[A-Za-z]+`)).
		Message("Valid City required").
		Key("userregister.City")
	v.Required(u.Zipcode).
		Message("Zipcode required").
		Key("userregister.Zipcode")
	v.Match(u.Zipcode, regexp.MustCompile("^([0-9]){5}$")).
		Message("Valid City required").
		Key("userregister.City")
	v.Required(u.Phnumber).
		Message("Phnumber required").
		Key("userregister.Phnumber")
	v.Match(u.Zipcode, regexp.MustCompile(`[0-9]{3}\-[0-9]{3}\-[0-9]{4}`)).
		Message("Phone Number must be in format DDD-DDD-DDDD").
		Key("userregister.Phnumber")
	v.Required(u.Email).
		Message("Email required").
		Key("userregister.Email")
	v.Email(u.Email).
		Message("Valid email required").
		Key("userregister.Email")
	v.Required(u.Password).
		Message("Password required").
		Key("userregister.Password")
	v.MinSize(u.Password, 8).
		Message("Password length must be at least 8").
		Key("userregister.Password")
	v.Required(u.PasswordConfirm == u.Password).
		Message("The passwords do not match.").
		Key("userregister.PasswordConfirm")
}
