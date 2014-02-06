package models

type UserAuthInterface interface {
	AuthId() string
	AuthSecret() string
}
