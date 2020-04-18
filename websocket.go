package phemex

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
)

// WsHandler handle raw websocket message
type WsHandler func(message *WsAOP)

// ErrHandler handles errors
type ErrHandler func(err error)

func keepAlive(c *websocket.Conn, id int64, stop chan struct{}, errHandler ErrHandler) {
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
			case <-ticker.C:
				id++
				p, _ := json.Marshal(map[string]interface{}{
					"id":     id,
					"method": "server.ping",
					"params": []string{},
				})
				c.WriteControl(websocket.PingMessage, p, time.Time{})
			case <-time.After(3*WebsocketTimeout - time.Since(lastResponse)):
				// Anything between 10s and 15s will come here
				errHandler(fmt.Errorf("last pong exceeded the timeout: %[1]v (%[2]v)", time.Since(lastResponse), id))
				return
			}
		}
	}()
}
