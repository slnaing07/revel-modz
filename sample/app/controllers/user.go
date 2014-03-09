package controllers

import (
	"github.com/revel/revel"

	"github.com/iassic/revel-modz/sample/app/routes"
)

type User struct {
	App
}

func (c User) CheckLoggedIn() revel.Result {
	if u := c.userConnected(); u == nil {
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

func (c User) Dashboard() revel.Result {
	return c.Render()
}

func (c User) FilesView() revel.Result {
	return c.Render()
}

// Defined in account.go
// func (c User) Account() revel.Result

// Defined in comm.go
// func (c User) Comm(ws *websocket.Conn) revel.Result
// func echoHandler(msg string, outbound chan string)
