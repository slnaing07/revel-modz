package ws_comm

import (
	"strings"

	"code.google.com/p/go.net/websocket"
)

// A function which takes arguements as a string
// and can send string replies via a channel
type WsHandler func(string, chan string)

type WsComm struct {
	handlers map[string]WsHandler
}

func New() *WsComm {
	h := new(WsComm)
	h.handlers = make(map[string]WsHandler)
	return h
}

// Add a message handler for a message type 'tag'
func (comm *WsComm) AddHandler(tag string, handle WsHandler) {
	comm.handlers[tag] = handle
}

func (comm *WsComm) RemoveHandler(tag string) {
	delete(comm.handlers, tag)
}

// Call 'go ws_comm.Serve(ws)' once the handlers are in place
func (comm *WsComm) Serve(ws *websocket.Conn) {

	// In order to select between inbound and outbound messages,
	// we need to stuff websocket events into a channel.
	inbound := make(chan string, 8)
	outbound := make(chan string, 8)

	// run the receive end of the socket in a goroutine
	go func() {
		var msg string
		for {
			err := websocket.Message.Receive(ws, &msg)
			if err != nil {
				close(inbound)
				return
			}
			inbound <- msg
		}
	}()

	// Now listen for new messages from the client or a handler
	for {
		select {
		case msg := <-outbound:
			if err := websocket.Message.Send(ws, msg); err != nil {
				// They disconnected.
				return
			}
		case msg, ok := <-inbound:
			// If the channel is closed, they disconnected.
			if !ok {
				return
			}

			// Otherwise, process message and possibly say something.
			// This spawns the handler in a goroutine, if found.
			comm.processMsg(msg, outbound)
		}
	}

	// websocket has closed
	return
}

func (comm *WsComm) processMsg(msg string, outChan chan string) {
	println("msg: " + msg)
	flds := strings.Fields(msg)
	tag := flds[0]

	handle, ok := comm.handlers[tag]
	if !ok {
		outChan <- "error unknown message type: '" + tag + "'"
		return
	}

	args := msg[len(tag):]

	go handle(args, outChan)

}
