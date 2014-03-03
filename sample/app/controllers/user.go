package controllers

import (

	"code.google.com/p/go.net/websocket"
	"github.com/iassic/revel-modz/modules/ws_comm"
	"github.com/revel/revel"

	"github.com/iassic/revel-modz/sample/app/routes"
)

type User struct {
	App
}

func (c User) CheckLoggedIn() revel.Result {
	if u := c.connected(); u == nil {
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

func (c User) Account() revel.Result {
	return c.Render()
}

func (c User) Dashboard() revel.Result {
	return c.Render()
}

func (c User) FilesView() revel.Result {
	return c.Render()
}

func (c User) Comm(ws *websocket.Conn) revel.Result {
	user := c.connected()
	revel.WARN.Println(user)

	comm := ws_comm.New()
	comm.AddHandler("echo", echoHandler)

	comm.Serve(ws)

	revel.INFO.Println("closing WS connection")
	return nil
}

func echoHandler(msg string, outbound chan string) {
	revel.INFO.Println("Echo: ", msg)
	outbound <- "ack " + msg
}
