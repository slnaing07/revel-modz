package controllers

import (
	// "fmt"

	"github.com/robfig/revel"

	// "github.com/iassic/revel-modz/samples/user_auth/app/routes"
)

type User struct {
	DbController
}

func (c User) Index() revel.Result {
	return c.Render()
}

func (c User) Result() revel.Result {
	return c.Render()
}
