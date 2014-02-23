package controllers

import (
	// "fmt"

	"github.com/robfig/revel"
	// "github.com/iassic/revel-modz/modules/user"

	"github.com/iassic/revel-modz/samples/user_auth/app/routes"
)

type User struct {
	App
}

func (c User) CheckLoggedIn() revel.Result {
	if u := c.connected(); u == nil {
		revel.ERROR.Println("Please log in first")
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
