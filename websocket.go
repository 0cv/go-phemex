package phemex

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
)

// WsHandler handle raw websocket message
type WsHandler func(message *WsAOP)

// ErrHandler handles errors
type ErrHandler func(err error)

func keepAlive(c *websocket.Conn, stop chan struct{}, errHandler ErrHandler) {
	rand.Seed(time.Now().UnixNano())
	ticker := time.NewTicker(WebsocketTimeout)

	lastResponse := time.Now()
	c.SetPongHandler(func(msg string) error {
		lastResponse = time.Now()
		return nil
	})

	go func() {
		defer func() {
			ticker.Stop()
		}()
		for {
			select {
			case <-stop:
				return
			default:
				deadline := time.Now().Add(10 * time.Second)
				c.WriteControl(websocket.PingMessage, []byte{}, deadline)

				<-ticker.C
				if time.Since(lastResponse) > WebsocketTimeout {
					errHandler(fmt.Errorf("last pong exceeded the timeout"))
					return
				}
			}
		}
	}()
}
