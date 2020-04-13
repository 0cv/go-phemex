package phemex

import (
	"context"
	"encoding/json"

	"github.com/Krisa/go-phemex/common"
)

// CreateOrderService create order
type CreateOrderService struct {
	c                *Client
	symbol           string
	clOrdID          *string
	actionBy         *string
	side             SideType
	orderQty         *float64
	priceEp          *int64
	ordType          *OrderType
	stopPxEp         *int64
	timeInForce      *TimeInForceType
	reduceOnly       *bool
	closeOnTrigger   *bool
	takeProfitEp     *int64
	stopLossEp       *int64
	pegOffsetValueEp *int64
	triggerType      *TriggerType
	text             *string
	pegPriceType     *string
}

// Symbol set symbol
func (s *CreateOrderService) Symbol(symbol string) *CreateOrderService {
	s.symbol = symbol
	return s
}

// ClOrdID set clOrID
func (s *CreateOrderService) ClOrdID(clOrdID string) *CreateOrderService {
	s.clOrdID = &clOrdID
	return s
}

// ActionBy set actionBy
func (s *CreateOrderService) ActionBy(actionBy string) *CreateOrderService {
	s.actionBy = &actionBy
	return s
}

// Side set side
func (s *CreateOrderService) Side(side SideType) *CreateOrderService {
	s.side = side
	return s
}

// OrderQty set orderQty
func (s *CreateOrderService) OrderQty(orderQty float64) *CreateOrderService {
	s.orderQty = &orderQty
	return s
}

// PriceEp set priceEp
func (s *CreateOrderService) PriceEp(priceEp int64) *CreateOrderService {
	s.priceEp = &priceEp
	return s
}

// OrdType set ordType
func (s *CreateOrderService) OrdType(ordType OrderType) *CreateOrderService {
	s.ordType = &ordType
	return s
}

// StopPxEp set stopPxEp
func (s *CreateOrderService) StopPxEp(stopPxEp int64) *CreateOrderService {
	s.stopPxEp = &stopPxEp
	return s
}

// TimeInForce set timeInForce
func (s *CreateOrderService) TimeInForce(timeInForce TimeInForceType) *CreateOrderService {
	s.timeInForce = &timeInForce
	return s
}

// ReduceOnly set reduceOnly
func (s *CreateOrderService) ReduceOnly(reduceOnly bool) *CreateOrderService {
	s.reduceOnly = &reduceOnly
	return s
}

// CloseOnTrigger set closeOnTrigger
func (s *CreateOrderService) CloseOnTrigger(closeOnTrigger bool) *CreateOrderService {
	s.closeOnTrigger = &closeOnTrigger
	return s
}

// TakeProfitEp set takeProfitEp
func (s *CreateOrderService) TakeProfitEp(takeProfitEp int64) *CreateOrderService {
	s.takeProfitEp = &takeProfitEp
	return s
}

// StopLossEp set stopLossEp
func (s *CreateOrderService) StopLossEp(stopLossEp int64) *CreateOrderService {
	s.stopLossEp = &stopLossEp
	return s
}

// TriggerType set triggerType
func (s *CreateOrderService) TriggerType(triggerType TriggerType) *CreateOrderService {
	s.triggerType = &triggerType
	return s
}

// Text set text
func (s *CreateOrderService) Text(text string) *CreateOrderService {
	s.text = &text
	return s
}

// PegOffsetValueEp set pegOffsetValueEp
func (s *CreateOrderService) PegOffsetValueEp(pegOffsetValueEp int64) *CreateOrderService {
	s.pegOffsetValueEp = &pegOffsetValueEp
	return s
}

// PegPriceType set pegPriceType
func (s *CreateOrderService) PegPriceType(pegPriceType string) *CreateOrderService {
	s.pegPriceType = &pegPriceType
	return s
}

func (s *CreateOrderService) createOrder(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   "POST",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"symbol": s.symbol,
		"side":   s.side,
	}
	if s.clOrdID != nil {
		m["clOrdID"] = *s.clOrdID
	}
	if s.orderQty != nil {
		m["orderQty"] = *s.orderQty
	}
	if s.actionBy != nil {
		m["actionBy"] = *s.actionBy
	}
	if s.priceEp != nil {
		m["priceEp"] = *s.priceEp
	}
	if s.ordType != nil {
		m["ordType"] = *s.ordType
	}
	if s.stopPxEp != nil {
		m["stopPxEp"] = *s.stopPxEp
	}
	if s.timeInForce != nil {
		m["timeInForce"] = *s.timeInForce
	}
	if s.reduceOnly != nil {
		m["reduceOnly"] = *s.reduceOnly
	}
	if s.closeOnTrigger != nil {
		m["closeOnTrigger"] = *s.closeOnTrigger
	}
	if s.takeProfitEp != nil {
		m["takeProfitEp"] = *s.takeProfitEp
	}
	if s.stopLossEp != nil {
		m["stopLossEp"] = *s.stopLossEp
	}
	if s.triggerType != nil {
		m["triggerType"] = *s.triggerType
	}
	if s.text != nil {
		m["text"] = *s.text
	}
	if s.pegPriceType != nil {
		m["pegPriceType"] = *s.pegPriceType
	}
	if s.pegOffsetValueEp != nil {
		m["pegOffsetValueEp"] = *s.pegOffsetValueEp
	}
	r.setFormParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *CreateOrderService) Do(ctx context.Context, opts ...RequestOption) (res *OrderResponse, err error) {
	data, err := s.createOrder(ctx, "/orders", opts...)
	if err != nil {
		return nil, err
	}
	resp := new(BaseResponse)
	resp.Data = new(OrderResponse)
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code > 0 {
		return nil, common.APIError{
			Code:    resp.Code,
			Message: resp.Msg,
		}
	}
	return resp.Data.(*OrderResponse), nil
}

// OrderResponse define create order response
type OrderResponse struct {
	BizError       int             `json:"bizError"`
	OrderID        string          `json:"orderID"`
	ClOrdID        string          `json:"clOrdID"`
	Symbol         string          `json:"symbol"`
	Side           SideType        `json:"side"`
	ActionTimeNs   int64           `json:"actionTimeNs"`
	TransactTimeNs int64           `json:"transactTimeNs"`
	OrderType      OrderType       `json:"orderType"`
	PriceEp        int64           `json:"priceEp"`
	Price          float64         `json:"price"`
	OrderQty       float64         `json:"orderQty"`
	DisplayQty     float64         `json:"displayQty"`
	TimeInForce    TimeInForceType `json:"timeInForce"`
	ReduceOnly     bool            `json:"reduceOnly"`
	TakeProfitEp   int64           `json:"takeProfitEp"`
	TakeProfit     float64         `json:"takeProfit"`
	StopPxEp       int64           `json:"stopPxEp"`
	StopPx         float64         `json:"stopPx"`
	StopLossEp     int64           `json:"stopLossEp"`
	ClosedPnlEv    int64           `json:"closedPnlEv"`
	ClosedPnl      float64         `json:"closedPnl"`
	ClosedSize     float64         `json:"closedSize"`
	CumQty         float64         `json:"cumQty"`
	CumValueEv     int64           `json:"cumValueEv"`
	CumValue       float64         `json:"cumValue"`
	LeavesQty      float64         `json:"leavesQty"`
	LeavesValueEv  int64           `json:"leavesValueEv"`
	LeavesValue    float64         `json:"leavesValue"`
	StopLoss       float64         `json:"stopLoss"`
	StopDirection  string          `json:"stopDirection"`
	OrdStatus      string          `json:"ordStatus"`
	Trigger        string          `json:"trigger"`
}

// CreateReplaceOrderService create order
type CreateReplaceOrderService struct {
	c            *Client
	symbol       string
	orderID      string
	origClOrdID  *string
	clOrdID      *string
	price        *float64
	priceEp      *int64
	orderQty     *float64
	stopPx       *float64
	stopPxEp     *int64
	takeProfit   *float64
	takeProfitEp *int64
	stopLoss     *float64
	stopLossEp   *int64
	pegOffset    *float64
	pegOffsetEp  *int64
}

// Symbol set symbol
func (s *CreateReplaceOrderService) Symbol(symbol string) *CreateReplaceOrderService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *CreateReplaceOrderService) OrderID(orderID string) *CreateReplaceOrderService {
	s.orderID = orderID
	return s
}

// OrigClOrdID set origClOrdID
func (s *CreateReplaceOrderService) OrigClOrdID(origClOrdID string) *CreateReplaceOrderService {
	s.origClOrdID = &origClOrdID
	return s
}

// ClOrdID set clOrID
func (s *CreateReplaceOrderService) ClOrdID(clOrdID string) *CreateReplaceOrderService {
	s.clOrdID = &clOrdID
	return s
}

// Price set price
func (s *CreateReplaceOrderService) Price(price float64) *CreateReplaceOrderService {
	s.price = &price
	return s
}

// PriceEp set priceEp
func (s *CreateReplaceOrderService) PriceEp(priceEp int64) *CreateReplaceOrderService {
	s.priceEp = &priceEp
	return s
}

// OrderQty set orderQty
func (s *CreateReplaceOrderService) OrderQty(orderQty float64) *CreateReplaceOrderService {
	s.orderQty = &orderQty
	return s
}

// StopPx set stopPx
func (s *CreateReplaceOrderService) StopPx(stopPx float64) *CreateReplaceOrderService {
	s.stopPx = &stopPx
	return s
}

// StopPxEp set stopPxEp
func (s *CreateReplaceOrderService) StopPxEp(stopPxEp int64) *CreateReplaceOrderService {
	s.stopPxEp = &stopPxEp
	return s
}

// TakeProfit set takeProfit
func (s *CreateReplaceOrderService) TakeProfit(takeProfit float64) *CreateReplaceOrderService {
	s.takeProfit = &takeProfit
	return s
}

// TakeProfitEp set takeProfitEp
func (s *CreateReplaceOrderService) TakeProfitEp(takeProfitEp int64) *CreateReplaceOrderService {
	s.takeProfitEp = &takeProfitEp
	return s
}

// StopLoss set stopLoss
func (s *CreateReplaceOrderService) StopLoss(stopLoss float64) *CreateReplaceOrderService {
	s.stopLoss = &stopLoss
	return s
}

// StopLossEp set stopLossEp
func (s *CreateReplaceOrderService) StopLossEp(stopLossEp int64) *CreateReplaceOrderService {
	s.stopLossEp = &stopLossEp
	return s
}

// PegOffset set pegOffset
func (s *CreateReplaceOrderService) PegOffset(pegOffset float64) *CreateReplaceOrderService {
	s.pegOffset = &pegOffset
	return s
}

// PegOffsetEp set pegOffsetEp
func (s *CreateReplaceOrderService) PegOffsetEp(pegOffsetEp int64) *CreateReplaceOrderService {
	s.pegOffsetEp = &pegOffsetEp
	return s
}

func (s *CreateReplaceOrderService) createOrder(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   "PUT",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	r.setParam("orderID", s.orderID)

	if s.origClOrdID != nil {
		r.setParam("origClOrdID", *s.origClOrdID)
	}
	if s.clOrdID != nil {
		r.setParam("clOrdID", *s.clOrdID)
	}
	if s.price != nil {
		r.setParam("price", *s.price)
	}
	if s.priceEp != nil {
		r.setParam("priceEp", *s.priceEp)
	}
	if s.orderQty != nil {
		r.setParam("orderQty", *s.orderQty)
	}
	if s.stopPx != nil {
		r.setParam("stopPx", *s.stopPx)
	}
	if s.stopPxEp != nil {
		r.setParam("stopPxEp", *s.stopPxEp)
	}
	if s.takeProfit != nil {
		r.setParam("takeProfit", *s.takeProfit)
	}
	if s.takeProfitEp != nil {
		r.setParam("takeProfitEp", *s.takeProfitEp)
	}
	if s.stopLoss != nil {
		r.setParam("stopLoss", *s.stopLoss)
	}
	if s.stopLossEp != nil {
		r.setParam("stopLossEp", *s.stopLossEp)
	}
	if s.pegOffset != nil {
		r.setParam("pegOffset", *s.pegOffset)
	}
	if s.pegOffsetEp != nil {
		r.setParam("pegOffsetEp", *s.pegOffsetEp)
	}
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *CreateReplaceOrderService) Do(ctx context.Context, opts ...RequestOption) (res *OrderResponse, err error) {
	data, err := s.createOrder(ctx, "/orders/replace", opts...)
	if err != nil {
		return nil, err
	}
	resp := new(BaseResponse)
	resp.Data = new(OrderResponse)
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code > 0 {
		return nil, common.APIError{
			Code:    resp.Code,
			Message: resp.Msg,
		}
	}
	return resp.Data.(*OrderResponse), nil
}

// ListOpenOrdersService list opened orders
type ListOpenOrdersService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *ListOpenOrdersService) Symbol(symbol string) *ListOpenOrdersService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *ListOpenOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []*OrderResponse, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/orders/activeList",
		secType:  secTypeSigned,
	}
	if s.symbol != "" {
		r.setParam("symbol", s.symbol)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*OrderResponse{}, err
	}

	resp := new(BaseResponse)
	resp.Data = new(RowsOrderResponse)

	err = json.Unmarshal(data, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code > 0 {
		return nil, common.APIError{
			Code:    resp.Code,
			Message: resp.Msg,
		}
	}
	rows := resp.Data.(*RowsOrderResponse)
	return rows.Rows, nil
}

// RowsOrderResponse rows order response
type RowsOrderResponse struct {
	Rows []*OrderResponse `json:"rows"`
}

// CancelOrderService cancel an order
type CancelOrderService struct {
	c       *Client
	symbol  string
	orderID *string
}

// Symbol set symbol
func (s *CancelOrderService) Symbol(symbol string) *CancelOrderService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *CancelOrderService) OrderID(orderID string) *CancelOrderService {
	s.orderID = &orderID
	return s
}

// Do send request
func (s *CancelOrderService) Do(ctx context.Context, opts ...RequestOption) (res *OrderResponse, err error) {
	r := &request{
		method:   "DELETE",
		endpoint: "/orders/cancel",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.orderID != nil {
		r.setParam("orderID", *s.orderID)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	resp := new(BaseResponse)
	resp.Data = new(OrderResponse)
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code > 0 {
		return nil, common.APIError{
			Code:    resp.Code,
			Message: resp.Msg,
		}
	}
	return resp.Data.(*OrderResponse), nil
}

// PositionsLeverageService cancel an order
type PositionsLeverageService struct {
	c          *Client
	symbol     string
	leverage   *int64
	leverageEr *int64
}

// Symbol set symbol
func (s *PositionsLeverageService) Symbol(symbol string) *PositionsLeverageService {
	s.symbol = symbol
	return s
}

// Leverage set leverage
func (s *PositionsLeverageService) Leverage(leverage int64) *PositionsLeverageService {
	s.leverage = &leverage
	return s
}

// LeverageEr set leverageEr
func (s *PositionsLeverageService) LeverageEr(leverageEr int64) *PositionsLeverageService {
	s.leverageEr = &leverageEr
	return s
}

// Do send request
func (s *PositionsLeverageService) Do(ctx context.Context, opts ...RequestOption) (res *BaseResponse, err error) {
	r := &request{
		method:   "PUT",
		endpoint: "/positions/leverage",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.leverage != nil {
		r.setParam("leverage", *s.leverage)
	}
	if s.leverageEr != nil {
		r.setParam("leverageEr", *s.leverageEr)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	resp := new(BaseResponse)
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code > 0 {
		return nil, common.APIError{
			Code:    resp.Code,
			Message: resp.Msg,
		}
	}
	return resp, nil
}

// PositionsAssignService cancel an order
type PositionsAssignService struct {
	c            *Client
	symbol       string
	posBalance   *float64
	posBalanceEv *int64
}

// Symbol set symbol
func (s *PositionsAssignService) Symbol(symbol string) *PositionsAssignService {
	s.symbol = symbol
	return s
}

// PosBalance set posBalance
func (s *PositionsAssignService) PosBalance(posBalance float64) *PositionsAssignService {
	s.posBalance = &posBalance
	return s
}

// PosBalanceEr set posBalanceEv
func (s *PositionsAssignService) PosBalanceEr(posBalanceEv int64) *PositionsAssignService {
	s.posBalanceEv = &posBalanceEv
	return s
}

// Do send request
func (s *PositionsAssignService) Do(ctx context.Context, opts ...RequestOption) (res *BaseResponse, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/positions/assign",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.posBalance != nil {
		r.setParam("posBalance", *s.posBalance)
	}
	if s.posBalanceEv != nil {
		r.setParam("posBalanceEv", *s.posBalanceEv)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	resp := new(BaseResponse)
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code > 0 {
		return nil, common.APIError{
			Code:    resp.Code,
			Message: resp.Msg,
		}
	}
	return resp, nil
}
