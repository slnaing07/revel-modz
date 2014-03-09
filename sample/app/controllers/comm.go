package controllers

import (
	"code.google.com/p/go.net/websocket"
	"github.com/iassic/revel-modz/modules/ws_comm"
	"github.com/revel/revel"
)

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
