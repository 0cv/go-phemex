package phemex

import (
	"context"
	"encoding/json"

	"github.com/Krisa/go-phemex/common"
)

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
		return nil, &common.APIError{
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
		return nil, &common.APIError{
			Code:    resp.Code,
			Message: resp.Msg,
		}
	}
	return resp, nil
}
