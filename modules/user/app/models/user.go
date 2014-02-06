package models

type UserBase struct {
	Identifier string
	Password   string `ignore:"yes"`
	HashPass   []byte
}

func (u UserBase) AuthId() string {
	return u.Indentifier
}
func (u UserBase) AuthSecret() string {
	return u.HashPass
}
