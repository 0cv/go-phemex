package phemex

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"strings"
)

// StartWsOrderBookService :
type StartWsOrderBookService struct {
	c       *Client
	id      *int64
	symbols []string
}

// ID :
func (s *StartWsOrderBookService) ID(id int64) *StartWsOrderBookService {
	s.id = &id
	return s
}

// Symbols :
func (s *StartWsOrderBookService) Symbols(symbols []string) *StartWsOrderBookService {
	s.symbols = symbols
	return s
}

// Do :
func (s *StartWsOrderBookService) Do(c *websocket.Conn, handler WsHandler, errHandler ErrHandler, opts ...RequestOption) (err error) {

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
			"method": "orderbook.subscribe",
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
				resp = new(WsOrderBookEvent)
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

// WsOrderBookEvent ws OrderBookEvent
type WsOrderBookEvent struct {
	Book      WsOrderBook `json:"book"`
	Depth     int         `json:"depth"`
	Sequence  int         `json:"sequence"`
	Symbol    string      `json:"symbol"`
	Timestamp int         `json:"timestamp"`
	Type      string      `json:"type"`
}

// WsOrderBook ws OrderBook
type WsOrderBook struct {
	Asks [][]int `json:"asks"`
	Bids [][]int `json:"bids"`
}
