package phemex

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type baseOrderTestSuite struct {
	baseTestSuite
}

type orderServiceTestSuite struct {
	baseOrderTestSuite
}

func TestOrderService(t *testing.T) {
	suite.Run(t, new(orderServiceTestSuite))
}

func (s *orderServiceTestSuite) TestOrderService() {
	data := []byte(`{
		"code": 0,
		"msg": "",
		"data": {
			"bizError": 0,
			"orderID": "ab90a08c-b728-4b6b-97c4-36fa497335bf",
			"clOrdID": "137e1928-5d25-fecd-dbd1-705ded659a4f",
			"symbol": "BTCUSD",
			"side": "Sell",
			"actionTimeNs": 1580547265848034600,
			"transactTimeNs": 0,
			"orderType": "Limit",
			"priceEp": 98970000,
			"price": 9897,
			"orderQty": 1,
			"displayQty": 1,
			"timeInForce": "GoodTillCancel",
			"reduceOnly": false,
			"stopPxEp": 0,
			"closedPnlEv": 0,
			"closedPnl": 0,
			"closedSize": 0,
			"cumQty": 0,
			"cumValueEv": 0,
			"cumValue": 0,
			"leavesQty": 1,
			"leavesValueEv": 10104,
			"leavesValue": 0.00010104,
			"stopPx": 0,
			"stopDirection": "UNSPECIFIED",
			"ordStatus": "Created"
		}
	}`)

	s.mockDo(data, nil)
	defer s.assertDo()
	symbol := "BTCUSD"
	side := SideTypeSell
	orderType := OrderTypeLimit
	timeInForce := TimeInForceTypeGTC
	price := int64(1000)
	quantity := float64(1000)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"ordType":     string(orderType),
			"orderQty":    quantity,
			"priceEp":     float64(price),
			"side":        string(side),
			"symbol":      symbol,
			"timeInForce": string(timeInForce),
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewCreateOrderService().Side(side).Symbol(symbol).OrdType(orderType).TimeInForce(TimeInForceTypeGTC).PriceEp(price).OrderQty(quantity).Do(newContext())
	s.r().NoError(err)
	e := &OrderResponse{
		BizError:       0,
		OrderID:        "ab90a08c-b728-4b6b-97c4-36fa497335bf",
		ClOrdID:        "137e1928-5d25-fecd-dbd1-705ded659a4f",
		Symbol:         "BTCUSD",
		Side:           side,
		ActionTimeNs:   1580547265848034600,
		TransactTimeNs: 0,
		OrderType:      orderType,
		PriceEp:        98970000,
		Price:          9897,
		OrderQty:       1,
		DisplayQty:     1,
		TimeInForce:    timeInForce,
		ReduceOnly:     false,
		StopPxEp:       0,
		ClosedPnlEv:    0,
		ClosedPnl:      0,
		ClosedSize:     0,
		CumQty:         0,
		CumValueEv:     0,
		CumValue:       0,
		LeavesQty:      1,
		LeavesValueEv:  10104,
		LeavesValue:    0.00010104,
		StopPx:         0,
		StopDirection:  "UNSPECIFIED",
		OrdStatus:      "Created",
	}
	s.assertOrderEquals(e, res)
}

func (s *orderServiceTestSuite) assertOrderEquals(e, a *OrderResponse) {
	r := s.r()
	r.Equal(e.BizError, a.BizError, "BizError")
	r.Equal(e.OrderID, a.OrderID, "OrderID")
	r.Equal(e.ClOrdID, a.ClOrdID, "ClOrdID")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Side, a.Side, "Side")
	r.Equal(e.ActionTimeNs, a.ActionTimeNs, "ActionTimeNs")
	r.Equal(e.TransactTimeNs, a.TransactTimeNs, "TransactTimeNs")
	r.Equal(e.OrderType, a.OrderType, "OrderType")
	r.Equal(e.PriceEp, a.PriceEp, "PriceEp")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.OrderQty, a.OrderQty, "OrderQty")
	r.Equal(e.DisplayQty, a.DisplayQty, "DisplayQty")
	r.Equal(e.TimeInForce, a.TimeInForce, "TimeInForce")
	r.Equal(e.ReduceOnly, a.ReduceOnly, "ReduceOnly")
	r.Equal(e.StopPxEp, a.StopPxEp, "StopPxEp")
	r.Equal(e.ClosedPnlEv, a.ClosedPnlEv, "ClosedPnlEv")
	r.Equal(e.ClosedPnl, a.ClosedPnl, "ClosedPnl")
	r.Equal(e.ClosedSize, a.ClosedSize, "ClosedSize")
	r.Equal(e.CumQty, a.CumQty, "CumQty")
	r.Equal(e.CumValueEv, a.CumValueEv, "CumValueEv")
	r.Equal(e.CumValue, a.CumValue, "CumValue")
	r.Equal(e.LeavesQty, a.LeavesQty, "LeavesQty")
	r.Equal(e.LeavesValueEv, a.LeavesValueEv, "LeavesValueEv")
	r.Equal(e.LeavesValue, a.LeavesValue, "LeavesValue")
	r.Equal(e.StopPx, a.StopPx, "StopPx")
	r.Equal(e.StopDirection, a.StopDirection, "StopDirection")
	r.Equal(e.OrdStatus, a.OrdStatus, "OrdStatus")
}

func (s *orderServiceTestSuite) TestReplaceOrderService() {
	data := []byte(`{
		"code": 0,
		"msg": "",
		"data": {
			"bizError": 0,
			"orderID": "ab90a08c-b728-4b6b-97c4-36fa497335bf",
			"clOrdID": "137e1928-5d25-fecd-dbd1-705ded659a4f",
			"symbol": "BTCUSD",
			"side": "Sell",
			"actionTimeNs": 1580547265848034600,
			"transactTimeNs": 0,
			"orderType": "Limit",
			"priceEp": 98970000,
			"price": 9897,
			"orderQty": 1,
			"displayQty": 1,
			"timeInForce": "GoodTillCancel",
			"reduceOnly": false,
			"stopPxEp": 0,
			"closedPnlEv": 0,
			"closedPnl": 0,
			"closedSize": 0,
			"cumQty": 0,
			"cumValueEv": 0,
			"cumValue": 0,
			"leavesQty": 1,
			"leavesValueEv": 10104,
			"leavesValue": 0.00010104,
			"stopPx": 0,
			"stopDirection": "UNSPECIFIED",
			"ordStatus": "Created"
		}
	}`)

	s.mockDo(data, nil)
	defer s.assertDo()
	symbol := "BTCUSD"
	side := SideTypeSell
	orderType := OrderTypeLimit
	timeInForce := TimeInForceTypeGTC
	price := int64(1000)
	quantity := float64(1000)
	orderID := "12"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"orderQty": quantity,
			"priceEp":  float64(price),
			"symbol":   symbol,
			"orderId":  orderID,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewCreateReplaceOrderService().Symbol(symbol).OrderID(orderID).PriceEp(price).OrderQty(quantity).Do(newContext())
	s.r().NoError(err)
	e := &OrderResponse{
		BizError:       0,
		OrderID:        "ab90a08c-b728-4b6b-97c4-36fa497335bf",
		ClOrdID:        "137e1928-5d25-fecd-dbd1-705ded659a4f",
		Symbol:         "BTCUSD",
		Side:           side,
		ActionTimeNs:   1580547265848034600,
		TransactTimeNs: 0,
		OrderType:      orderType,
		PriceEp:        98970000,
		Price:          9897,
		OrderQty:       1,
		DisplayQty:     1,
		TimeInForce:    timeInForce,
		ReduceOnly:     false,
		StopPxEp:       0,
		ClosedPnlEv:    0,
		ClosedPnl:      0,
		ClosedSize:     0,
		CumQty:         0,
		CumValueEv:     0,
		CumValue:       0,
		LeavesQty:      1,
		LeavesValueEv:  10104,
		LeavesValue:    0.00010104,
		StopPx:         0,
		StopDirection:  "UNSPECIFIED",
		OrdStatus:      "Created",
	}
	s.assertReplaceOrderEquals(e, res)
}

func (s *orderServiceTestSuite) assertReplaceOrderEquals(e, a *OrderResponse) {
	r := s.r()
	r.Equal(e.BizError, a.BizError, "BizError")
	r.Equal(e.OrderID, a.OrderID, "OrderID")
	r.Equal(e.ClOrdID, a.ClOrdID, "ClOrdID")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Side, a.Side, "Side")
	r.Equal(e.ActionTimeNs, a.ActionTimeNs, "ActionTimeNs")
	r.Equal(e.TransactTimeNs, a.TransactTimeNs, "TransactTimeNs")
	r.Equal(e.OrderType, a.OrderType, "OrderType")
	r.Equal(e.PriceEp, a.PriceEp, "PriceEp")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.OrderQty, a.OrderQty, "OrderQty")
	r.Equal(e.DisplayQty, a.DisplayQty, "DisplayQty")
	r.Equal(e.TimeInForce, a.TimeInForce, "TimeInForce")
	r.Equal(e.ReduceOnly, a.ReduceOnly, "ReduceOnly")
	r.Equal(e.StopPxEp, a.StopPxEp, "StopPxEp")
	r.Equal(e.ClosedPnlEv, a.ClosedPnlEv, "ClosedPnlEv")
	r.Equal(e.ClosedPnl, a.ClosedPnl, "ClosedPnl")
	r.Equal(e.ClosedSize, a.ClosedSize, "ClosedSize")
	r.Equal(e.CumQty, a.CumQty, "CumQty")
	r.Equal(e.CumValueEv, a.CumValueEv, "CumValueEv")
	r.Equal(e.CumValue, a.CumValue, "CumValue")
	r.Equal(e.LeavesQty, a.LeavesQty, "LeavesQty")
	r.Equal(e.LeavesValueEv, a.LeavesValueEv, "LeavesValueEv")
	r.Equal(e.LeavesValue, a.LeavesValue, "LeavesValue")
	r.Equal(e.StopPx, a.StopPx, "StopPx")
	r.Equal(e.StopDirection, a.StopDirection, "StopDirection")
	r.Equal(e.OrdStatus, a.OrdStatus, "OrdStatus")
}

func (s *orderServiceTestSuite) TestListOpenService() {
	data := []byte(`{
		"code": 0,
		"msg": "",
		"data": {
			"rows": [{
				"bizError": 0,
				"orderID": "ab90a08c-b728-4b6b-97c4-36fa497335bf",
				"clOrdID": "137e1928-5d25-fecd-dbd1-705ded659a4f",
				"symbol": "BTCUSD",
				"side": "Sell",
				"actionTimeNs": 1580547265848034600,
				"transactTimeNs": 0,
				"orderType": "Limit",
				"priceEp": 98970000,
				"price": 9897,
				"orderQty": 1,
				"displayQty": 1,
				"timeInForce": "GoodTillCancel",
				"reduceOnly": false,
				"stopPxEp": 0,
				"closedPnlEv": 0,
				"closedPnl": 0,
				"closedSize": 0,
				"cumQty": 0,
				"cumValueEv": 0,
				"cumValue": 0,
				"leavesQty": 1,
				"leavesValueEv": 10104,
				"leavesValue": 0.00010104,
				"stopPx": 0,
				"stopDirection": "UNSPECIFIED",
				"ordStatus": "Created"
			}]
		}
	}`)

	s.mockDo(data, nil)
	defer s.assertDo()
	symbol := "BTCUSD"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol": symbol,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewListOpenOrdersService().Symbol(symbol).Do(newContext())
	s.r().NoError(err)
	e := []*OrderResponse{{
		BizError:       0,
		OrderID:        "ab90a08c-b728-4b6b-97c4-36fa497335bf",
		ClOrdID:        "137e1928-5d25-fecd-dbd1-705ded659a4f",
		Symbol:         "BTCUSD",
		Side:           "Sell",
		ActionTimeNs:   1580547265848034600,
		TransactTimeNs: 0,
		OrderType:      "Limit",
		PriceEp:        98970000,
		Price:          9897,
		OrderQty:       1,
		DisplayQty:     1,
		TimeInForce:    "GoodTillCancel",
		ReduceOnly:     false,
		StopPxEp:       0,
		ClosedPnlEv:    0,
		ClosedPnl:      0,
		ClosedSize:     0,
		CumQty:         0,
		CumValueEv:     0,
		CumValue:       0,
		LeavesQty:      1,
		LeavesValueEv:  10104,
		LeavesValue:    0.00010104,
		StopPx:         0,
		StopDirection:  "UNSPECIFIED",
		OrdStatus:      "Created",
	}}
	s.assertListOpenEquals(e, res)
}

func (s *orderServiceTestSuite) assertListOpenEquals(e, a []*OrderResponse) {
	r := s.r()
	for i := 0; i < len(a); i++ {
		r.Equal(e[i].BizError, a[i].BizError, "BizError")
		r.Equal(e[i].OrderID, a[i].OrderID, "OrderID")
		r.Equal(e[i].ClOrdID, a[i].ClOrdID, "ClOrdID")
		r.Equal(e[i].Symbol, a[i].Symbol, "Symbol")
		r.Equal(e[i].Side, a[i].Side, "Side")
		r.Equal(e[i].ActionTimeNs, a[i].ActionTimeNs, "ActionTimeNs")
		r.Equal(e[i].TransactTimeNs, a[i].TransactTimeNs, "TransactTimeNs")
		r.Equal(e[i].OrderType, a[i].OrderType, "OrderType")
		r.Equal(e[i].PriceEp, a[i].PriceEp, "PriceEp")
		r.Equal(e[i].Price, a[i].Price, "Price")
		r.Equal(e[i].OrderQty, a[i].OrderQty, "OrderQty")
		r.Equal(e[i].DisplayQty, a[i].DisplayQty, "DisplayQty")
		r.Equal(e[i].TimeInForce, a[i].TimeInForce, "TimeInForce")
		r.Equal(e[i].ReduceOnly, a[i].ReduceOnly, "ReduceOnly")
		r.Equal(e[i].StopPxEp, a[i].StopPxEp, "StopPxEp")
		r.Equal(e[i].ClosedPnlEv, a[i].ClosedPnlEv, "ClosedPnlEv")
		r.Equal(e[i].ClosedPnl, a[i].ClosedPnl, "ClosedPnl")
		r.Equal(e[i].ClosedSize, a[i].ClosedSize, "ClosedSize")
		r.Equal(e[i].CumQty, a[i].CumQty, "CumQty")
		r.Equal(e[i].CumValueEv, a[i].CumValueEv, "CumValueEv")
		r.Equal(e[i].CumValue, a[i].CumValue, "CumValue")
		r.Equal(e[i].LeavesQty, a[i].LeavesQty, "LeavesQty")
		r.Equal(e[i].LeavesValueEv, a[i].LeavesValueEv, "LeavesValueEv")
		r.Equal(e[i].LeavesValue, a[i].LeavesValue, "LeavesValue")
		r.Equal(e[i].StopPx, a[i].StopPx, "StopPx")
		r.Equal(e[i].StopDirection, a[i].StopDirection, "StopDirection")
		r.Equal(e[i].OrdStatus, a[i].OrdStatus, "OrdStatus")
	}
}

func (s *orderServiceTestSuite) TestDeleteOrderService() {
	data := []byte(`{
		"code": 0,
		"msg": "",
		"data": {
			"bizError": 0,
			"orderID": "ab90a08c-b728-4b6b-97c4-36fa497335bf",
			"clOrdID": "137e1928-5d25-fecd-dbd1-705ded659a4f",
			"symbol": "BTCUSD",
			"side": "Sell",
			"actionTimeNs": 1580547265848034600,
			"transactTimeNs": 0,
			"orderType": "Limit",
			"priceEp": 98970000,
			"price": 9897,
			"orderQty": 1,
			"displayQty": 1,
			"timeInForce": "GoodTillCancel",
			"reduceOnly": false,
			"stopPxEp": 0,
			"closedPnlEv": 0,
			"closedPnl": 0,
			"closedSize": 0,
			"cumQty": 0,
			"cumValueEv": 0,
			"cumValue": 0,
			"leavesQty": 1,
			"leavesValueEv": 10104,
			"leavesValue": 0.00010104,
			"stopPx": 0,
			"stopDirection": "UNSPECIFIED",
			"ordStatus": "Created"
		}
	}`)

	s.mockDo(data, nil)
	defer s.assertDo()
	symbol := "BTCUSD"
	orderID := "12"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":  symbol,
			"orderId": orderID,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewCancelOrderService().Symbol(symbol).OrderID(orderID).Do(newContext())
	s.r().NoError(err)
	e := &OrderResponse{
		BizError:       0,
		OrderID:        "ab90a08c-b728-4b6b-97c4-36fa497335bf",
		ClOrdID:        "137e1928-5d25-fecd-dbd1-705ded659a4f",
		Symbol:         "BTCUSD",
		Side:           "Sell",
		ActionTimeNs:   1580547265848034600,
		TransactTimeNs: 0,
		OrderType:      "Limit",
		PriceEp:        98970000,
		Price:          9897,
		OrderQty:       1,
		DisplayQty:     1,
		TimeInForce:    "GoodTillCancel",
		ReduceOnly:     false,
		StopPxEp:       0,
		ClosedPnlEv:    0,
		ClosedPnl:      0,
		ClosedSize:     0,
		CumQty:         0,
		CumValueEv:     0,
		CumValue:       0,
		LeavesQty:      1,
		LeavesValueEv:  10104,
		LeavesValue:    0.00010104,
		StopPx:         0,
		StopDirection:  "UNSPECIFIED",
		OrdStatus:      "Created",
	}
	s.assertDeleteOrderEquals(e, res)
}

func (s *orderServiceTestSuite) assertDeleteOrderEquals(e, a *OrderResponse) {
	r := s.r()
	r.Equal(e.BizError, a.BizError, "BizError")
	r.Equal(e.OrderID, a.OrderID, "OrderID")
	r.Equal(e.ClOrdID, a.ClOrdID, "ClOrdID")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Side, a.Side, "Side")
	r.Equal(e.ActionTimeNs, a.ActionTimeNs, "ActionTimeNs")
	r.Equal(e.TransactTimeNs, a.TransactTimeNs, "TransactTimeNs")
	r.Equal(e.OrderType, a.OrderType, "OrderType")
	r.Equal(e.PriceEp, a.PriceEp, "PriceEp")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.OrderQty, a.OrderQty, "OrderQty")
	r.Equal(e.DisplayQty, a.DisplayQty, "DisplayQty")
	r.Equal(e.TimeInForce, a.TimeInForce, "TimeInForce")
	r.Equal(e.ReduceOnly, a.ReduceOnly, "ReduceOnly")
	r.Equal(e.StopPxEp, a.StopPxEp, "StopPxEp")
	r.Equal(e.ClosedPnlEv, a.ClosedPnlEv, "ClosedPnlEv")
	r.Equal(e.ClosedPnl, a.ClosedPnl, "ClosedPnl")
	r.Equal(e.ClosedSize, a.ClosedSize, "ClosedSize")
	r.Equal(e.CumQty, a.CumQty, "CumQty")
	r.Equal(e.CumValueEv, a.CumValueEv, "CumValueEv")
	r.Equal(e.CumValue, a.CumValue, "CumValue")
	r.Equal(e.LeavesQty, a.LeavesQty, "LeavesQty")
	r.Equal(e.LeavesValueEv, a.LeavesValueEv, "LeavesValueEv")
	r.Equal(e.LeavesValue, a.LeavesValue, "LeavesValue")
	r.Equal(e.StopPx, a.StopPx, "StopPx")
	r.Equal(e.StopDirection, a.StopDirection, "StopDirection")
	r.Equal(e.OrdStatus, a.OrdStatus, "OrdStatus")
}

func (s *orderServiceTestSuite) TestQueryService() {
	data := []byte(`{
		"code": 0,
		"msg": "",
		"data": [{
			"bizError": 0,
			"orderID": "ab90a08c-b728-4b6b-97c4-36fa497335bf",
			"clOrdID": "137e1928-5d25-fecd-dbd1-705ded659a4f",
			"symbol": "BTCUSD",
			"side": "Sell",
			"actionTimeNs": 1580547265848034600,
			"transactTimeNs": 0,
			"orderType": "Limit",
			"priceEp": 98970000,
			"price": 9897,
			"orderQty": 1,
			"displayQty": 1,
			"timeInForce": "GoodTillCancel",
			"reduceOnly": false,
			"stopPxEp": 0,
			"closedPnlEv": 0,
			"closedPnl": 0,
			"closedSize": 0,
			"cumQty": 0,
			"cumValueEv": 0,
			"cumValue": 0,
			"leavesQty": 1,
			"leavesValueEv": 10104,
			"leavesValue": 0.00010104,
			"stopPx": 0,
			"stopDirection": "UNSPECIFIED",
			"ordStatus": "Created"
		}]
	}`)

	s.mockDo(data, nil)
	defer s.assertDo()
	symbol := "BTCUSD"
	orderID := "1234"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":  symbol,
			"orderId": orderID,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewQueryOrderService().Symbol(symbol).OrderID(orderID).Do(newContext())
	s.r().NoError(err)
	e := []*OrderResponse{{
		BizError:       0,
		OrderID:        "ab90a08c-b728-4b6b-97c4-36fa497335bf",
		ClOrdID:        "137e1928-5d25-fecd-dbd1-705ded659a4f",
		Symbol:         "BTCUSD",
		Side:           "Sell",
		ActionTimeNs:   1580547265848034600,
		TransactTimeNs: 0,
		OrderType:      "Limit",
		PriceEp:        98970000,
		Price:          9897,
		OrderQty:       1,
		DisplayQty:     1,
		TimeInForce:    "GoodTillCancel",
		ReduceOnly:     false,
		StopPxEp:       0,
		ClosedPnlEv:    0,
		ClosedPnl:      0,
		ClosedSize:     0,
		CumQty:         0,
		CumValueEv:     0,
		CumValue:       0,
		LeavesQty:      1,
		LeavesValueEv:  10104,
		LeavesValue:    0.00010104,
		StopPx:         0,
		StopDirection:  "UNSPECIFIED",
		OrdStatus:      "Created",
	}}
	s.assertListOpenEquals(e, res)
}
