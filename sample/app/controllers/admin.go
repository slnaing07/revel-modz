package controllers

import (
	"github.com/revel/revel"
)

type Admin struct {
	User
}

func (c Admin) Index() revel.Result {
	return c.Render()
}

func (c Admin) MaillistView() revel.Result {
	return c.Render()
}

func (c Admin) ComposeMaillistMessage() revel.Result {
	return c.Render()
}
