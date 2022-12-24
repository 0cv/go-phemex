package phemex

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"strings"
)

// StartWsTradeService :
type StartWsTradeService struct {
	c       *Client
	id      *int64
	symbols []string
}

// ID :
func (s *StartWsTradeService) ID(id int64) *StartWsTradeService {
	s.id = &id
	return s
}

// Symbols :
func (s *StartWsTradeService) Symbols(symbols []string) *StartWsTradeService {
	s.symbols = symbols
	return s
}

// Do :
func (s *StartWsTradeService) Do(c *websocket.Conn, handler WsHandler, errHandler ErrHandler, opts ...RequestOption) (err error) {

	stop := make(chan struct{})

	if c == nil {
		return fmt.Errorf("the connection is not initialized (%v)", *s.id)
	}
	go func() {
		defer close(stop)

		if WebsocketKeepalive {
			keepAlive(c, *s.id, stop, errHandler)
		}

		err = c.WriteJSON(map[string]interface{}{
			"id":     *s.id,
			"method": "trade.subscribe",
			"params": s.symbols,
		})

		if err != nil {
			errHandler(err)
			return
		}

		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				errHandler(err)
				return
			}
			var resp interface{}
			switch {
			case strings.HasPrefix(string(msg), "{\"error\""):
				resp = new(WsError)
			default:
				resp = new(WsTradeEvent)
			}
			err = json.Unmarshal(msg, &resp)

			if err != nil {
				errHandler(err)
				return
			}
			handler(resp)
		}
	}()
	return
}

// WsTradeEvent ws WsTradeEvent
type WsTradeEvent struct {
	Trade     []WsTrade `json:"trades"`
	Sequence  int       `json:"sequence"`
	Symbol    string    `json:"symbol"`
	Timestamp int       `json:"timestamp"`
	Type      string    `json:"type"`
}

type WsTrade struct {
}
