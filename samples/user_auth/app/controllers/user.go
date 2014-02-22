package controllers

import (
	// "fmt"

	"github.com/robfig/revel"

	"github.com/iassic/revel-modz/samples/user_auth/app/routes"
	// "github.com/iassic/revel-modz/modules/user"
)

type User struct {
	App
}

func (c User) CheckLoggedIn() revel.Result {
	if u := c.connected(); u != nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(routes.App.Login())
	}
	return nil
}

func (c User) Index() revel.Result {
	return c.Render()
}

func (c User) Result() revel.Result {
	return c.Render()
}
