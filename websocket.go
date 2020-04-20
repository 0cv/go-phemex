package phemex

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
)

// WsHandler handle raw websocket message
type WsHandler func(message interface{})

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

				// as suggested https://github.com/phemex/phemex-api-docs/blob/master/Public-API-en.md#1-session-management
				if time.Since(lastResponse) > 3*WebsocketTimeout {
					errHandler(fmt.Errorf("last pong exceeded the timeout: %[1]v (%[2]v)", time.Since(lastResponse), id))
					return
				}
			}
		}
	}()
}
