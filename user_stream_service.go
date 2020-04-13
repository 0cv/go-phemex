package phemex

import (
	"context"
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

var (
	baseURL = "wss://phemex.com/ws"
	// WebsocketTimeout is an interval for sending ping/pong messages if WebsocketKeepalive is enabled
	WebsocketTimeout = 5 * time.Second
	// WebsocketKeepalive enables sending ping/pong messages to check the connection stability
	WebsocketKeepalive = true
)

// WsAuthService create listen key for user stream service
type WsAuthService struct {
	c   *Client
	url *string
}

// URL set url. wss://testnet.phemex.com/ws for test net
func (s *WsAuthService) URL(url string) *WsAuthService {
	s.url = &url
	return s
}

// Do send request
func (s *WsAuthService) Do(ctx context.Context, opts ...RequestOption) (c *websocket.Conn, err error) {
	if s.url != nil {
		baseURL = *s.url
	}
	s.c.debug("dial URL: %s", baseURL)

	c, _, err = websocket.DefaultDialer.Dial(baseURL, nil)
	//c.SetReadLimit()
	if err != nil {
		return nil, err
	}

	expiry := currentTimestamp() + 60
	raw := fmt.Sprintf("%s%v", s.c.APIKey, expiry)
	signedString, err := s.c.signString(raw)
	if err != nil {
		return nil, err
	}

	err = c.WriteJSON(map[string]interface{}{
		"method": "user.auth",
		"params": []interface{}{
			"API",
			s.c.APIKey,
			signedString,
			expiry,
		},
		"id": 100,
	})

	if err != nil {
		return nil, err
	}

	_, _, err = c.ReadMessage()

	if err != nil {
		return nil, err
	}

	return c, nil
}

// StartWsAOPService create listen key for user stream service
type StartWsAOPService struct {
	c  *Client
	id *int
}

// SetID set id
func (s *StartWsAOPService) SetID(id int) *StartWsAOPService {
	s.id = &id
	return s
}

// Do send request
func (s *StartWsAOPService) Do(c *websocket.Conn, handler WsHandler, errHandler ErrHandler, opts ...RequestOption) (err error) {

	stop := make(chan struct{})

	if c == nil {
		return fmt.Errorf("the connection is not initialized (%v)", *s.id)
	}
	go func() {
		defer close(stop)

		if WebsocketKeepalive {
			keepAlive(c, stop, errHandler)
		}

		err = c.WriteJSON(map[string]interface{}{
			"id":     *s.id,
			"method": "aop.subscribe",
			"params": []interface{}{},
		})

		if err != nil {
			errHandler(err)
			return
		}

		for {
			resp := new(WsAOP)
			err := c.ReadJSON(resp)

			if err != nil {
				errHandler(err)
				return
			}
			handler(resp)
		}
	}()
	return
}

// WsAccount ws account
type WsAccount struct {
	AccountBalanceEv   int64  `json:"accountBalanceEv"`
	AccountID          int64  `json:"accountID"`
	BonusBalanceEv     int64  `json:"bonusBalanceEv"`
	Currency           string `json:"currency"`
	TotalUsedBalanceEv int64  `json:"totalUsedBalanceEv"`
	UserID             int64  `json:"userID"`
}

// WsOrder ws order
type WsOrder struct {
	AccountID               int64   `json:"accountID"`
	Action                  string  `json:"action"`
	ActionBy                string  `json:"actionBy"`
	ActionTimeNs            int64   `json:"actionTimeNs"`
	AddedSeq                int64   `json:"addedSeq"`
	BonusChangedAmountEv    int64   `json:"bonusChangedAmountEv"`
	ClOrdID                 string  `json:"clOrdID"`
	ClosedPnlEv             int64   `json:"closedPnlEv"`
	ClosedSize              float64 `json:"closedSize"`
	Code                    int64   `json:"code"`
	CumQty                  float64 `json:"cumQty"`
	CurAccBalanceEv         int64   `json:"curAccBalanceEv"`
	CurAssignedPosBalanceEv int64   `json:"curAssignedPosBalanceEv"`
	CurLeverageEr           int64   `json:"curLeverageEr"`
	CurPosSide              string  `json:"curPosSide"`
	CurPosSize              float64 `json:"curPosSize"`
	CurPosTerm              int64   `json:"curPosTerm"`
	CurPosValueEv           int64   `json:"curPosValueEv"`
	CurRiskLimitEv          int64   `json:"curRiskLimitEv"`
	Currency                string  `json:"currency"`
	CxlRejReason            int64   `json:"cxlRejReason"`
	DisplayQty              float64 `json:"displayQty"`
	ExecFeeEv               int64   `json:"execFeeEv"`
	ExecID                  string  `json:"execID"`
	ExecInst                string  `json:"execInst"`
	ExecPriceEp             int64   `json:"execPriceEp"`
	ExecQty                 float64 `json:"execQty"`
	ExecSeq                 float64 `json:"execSeq"`
	ExecStatus              string  `json:"execStatus"`
	ExecValueEv             int64   `json:"execValueEv"`
	FeeRateEr               int64   `json:"feeRateEr"`
	LeavesQty               float64 `json:"leavesQty"`
	LeavesValueEv           int64   `json:"leavesValueEv"`
	Message                 string  `json:"message"`
	OrdStatus               string  `json:"ordStatus"`
	OrdType                 string  `json:"ordType"`
	OrderID                 string  `json:"orderID"`
	OrderQty                float64 `json:"orderQty"`
	PegOffsetValueEp        int64   `json:"pegOffsetValueEp"`
	Platform                string  `json:"platform"`
	PriceEp                 int64   `json:"priceEp"`
	RelatedPosTerm          int64   `json:"relatedPosTerm"`
	RelatedReqNum           int64   `json:"relatedReqNum"`
	Side                    string  `json:"side"`
	StopLossEp              int64   `json:"stopLossEp"`
	StopPxEp                int64   `json:"stopPxEp"`
	Symbol                  string  `json:"symbol"`
	TakeProfitEp            int64   `json:"takeProfitEp"`
	TimeInForce             string  `json:"timeInForce"`
	TransactTimeNs          int64   `json:"transactTimeNs"`
	UserID                  int64   `json:"userID"`
}

// WsPosition ws position
type WsPosition struct {
	AccountID              int64   `json:"accountID"`
	ActionTimeNs           int64   `json:"actionTimeNs"`
	AssignedPosBalanceEv   int64   `json:"assignedPosBalanceEv"`
	AvgEntryPriceEp        int64   `json:"avgEntryPriceEp"`
	BankruptCommEv         int64   `json:"bankruptCommEv"`
	BankruptPriceEp        int64   `json:"bankruptPriceEp"`
	BuyLeavesQty           float64 `json:"buyLeavesQty"`
	BuyLeavesValueEv       int64   `json:"buyLeavesValueEv"`
	BuyValueToCostEr       int64   `json:"buyValueToCostEr"`
	CreatedAtNs            int64   `json:"createdAtNs"`
	CrossSharedBalanceEv   int64   `json:"crossSharedBalanceEv"`
	CumClosedPnlEv         int64   `json:"cumClosedPnlEv"`
	CumFundingFeeEv        int64   `json:"cumFundingFeeEv"`
	CumTransactFeeEv       int64   `json:"cumTransactFeeEv"`
	CurTermRealisedPnlEv   int64   `json:"curTermRealisedPnlEv"`
	Currency               string  `json:"currency"`
	DataVer                float64 `json:"dataVer"`
	DeleveragePercentileEr int64   `json:"deleveragePercentileEr"`
	DisplayLeverageEr      int64   `json:"displayLeverageEr"`
	EstimatedOrdLossEv     int64   `json:"estimatedOrdLossEv"`
	ExecSeq                int64   `json:"execSeq"`
	FreeCostEv             int64   `json:"freeCostEv"`
	FreeQty                float64 `json:"freeQty"`
	InitMarginReqEr        int64   `json:"initMarginReqEr"`
	LastFundingTime        int64   `json:"lastFundingTime"`
	LastTermEndTime        int64   `json:"lastTermEndTime"`
	LeverageEr             int64   `json:"leverageEr"`
	LiquidationPriceEp     int64   `json:"liquidationPriceEp"`
	MaintMarginReqEr       int64   `json:"maintMarginReqEr"`
	MakerFeeRateEr         int64   `json:"makerFeeRateEr"`
	MarkPriceEp            int64   `json:"markPriceEp"`
	OrderCostEv            int64   `json:"orderCostEv"`
	PosCostEv              int64   `json:"posCostEv"`
	PositionMarginEv       int64   `json:"positionMarginEv"`
	PositionStatus         string  `json:"positionStatus"`
	RiskLimitEv            int64   `json:"riskLimitEv"`
	SellLeavesQty          float64 `json:"sellLeavesQty"`
	SellLeavesValueEv      int64   `json:"sellLeavesValueEv"`
	SellValueToCostEr      int64   `json:"sellValueToCostEr"`
	Side                   string  `json:"side"`
	Size                   float64 `json:"size"`
	Symbol                 string  `json:"symbol"`
	TakerFeeRateEr         int64   `json:"takerFeeRateEr"`
	Term                   int64   `json:"term"`
	UnrealisedPnlEv        int64   `json:"unrealisedPnlEv"`
	UpdatedAtNs            int64   `json:"updatedAtNs"`
	UsedBalanceEv          int64   `json:"usedBalanceEv"`
	UserID                 int64   `json:"userID"`
	ValueEv                int64   `json:"valueEv"`
}

// WsPositionInfo ws position info
type WsPositionInfo struct {
	AccountID int64   `json:"accountID"`
	Light     float64 `json:"light"`
	Symbol    string  `json:"symbol"`
	UserID    int64   `json:"userID"`
}

// WsAOP ws AOP
type WsAOP struct {
	Accounts     []*WsAccount    `json:"accounts"`
	Orders       []*WsOrder      `json:"orders"`
	Positions    []*WsPosition   `json:"positions"`
	PositionInfo *WsPositionInfo `json:"position_info"`
	Sequence     int64           `json:"sequence"`
	Timestamp    int64           `json:"timestamp"`
	Type         string          `json:"type"`
}
