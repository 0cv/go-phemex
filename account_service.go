package phemex

import (
	"context"
	"encoding/json"
)

// GetAccountPositionService get account info
type GetAccountPositionService struct {
	c        *Client
	currency string
}

// Currency set currency
func (s *GetAccountPositionService) Currency(currency string) *GetAccountPositionService {
	s.currency = currency
	return s
}

// Do send request
func (s *GetAccountPositionService) Do(ctx context.Context, opts ...RequestOption) (*AccountPosition, error) {
	r := &request{
		method:   "GET",
		endpoint: "/accounts/accountPositions",
		secType:  secTypeSigned,
	}

	r.setParam("currency", s.currency)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	resp := new(BaseResponse)
	resp.Data = new(AccountPosition)
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	return resp.Data.(*AccountPosition), nil
}

// AccountPosition define account position info
type AccountPosition struct {
	Account   Account     `json:"account"`
	Positions []*Position `json:"positions"`
}

// Account account detail
type Account struct {
	AccountID          int64   `json:"accountId"`
	Currency           string  `json:"currency"`
	AccountBalanceEv   float64 `json:"accountBalanceEv"`
	TotalUsedBalanceEv float64 `json:"totalUsedBalanceEv"`
}

// Position position detail
type Position struct {
	AccountID              int64   `json:"accountID"`
	Symbol                 string  `json:"symbol"`
	Currency               string  `json:"currency"`
	Side                   string  `json:"side"`
	PositionStatus         string  `json:"positionStatus"`
	CrossMargin            bool    `json:"crossMargin"`
	LeverageEr             float64 `json:"leverageEr"`
	Leverage               float64 `json:"leverage"`
	InitMarginReqEr        float64 `json:"initMarginReqEr"`
	InitMarginReq          float64 `json:"initMarginReq"`
	MaintMarginReqEr       float64 `json:"maintMarginReqEr"`
	MaintMarginReq         float64 `json:"maintMarginReq"`
	RiskLimitEv            float64 `json:"riskLimitEv"`
	RiskLimit              float64 `json:"riskLimit"`
	Size                   float64 `json:"size"`
	Value                  float64 `json:"value"`
	ValueEv                float64 `json:"valueEv"`
	AvgEntryPriceEp        float64 `json:"avgEntryPriceEp"`
	AvgEntryPrice          float64 `json:"avgEntryPrice"`
	PosCostEv              float64 `json:"posCostEv"`
	PosCost                float64 `json:"posCost"`
	AssignedPosBalanceEv   float64 `json:"assignedPosBalanceEv"`
	AssignedPosBalance     float64 `json:"assignedPosBalance"`
	BankruptCommEv         float64 `json:"bankruptCommEv"`
	BankruptComm           float64 `json:"bankruptComm"`
	BankruptPriceEp        float64 `json:"bankruptPriceEp"`
	BankruptPrice          float64 `json:"bankruptPrice"`
	PositionMarginEv       float64 `json:"positionMarginEv"`
	PositionMargin         float64 `json:"positionMargin"`
	LiquidationPriceEp     float64 `json:"liquidationPriceEp"`
	LiquidationPrice       float64 `json:"liquidationPrice"`
	DeleveragePercentileEr float64 `json:"deleveragePercentileEr"`
	DeleveragePercentile   float64 `json:"deleveragePercentile"`
	BuyValueToCostEr       float64 `json:"buyValueToCostEr"`
	BuyValueToCost         float64 `json:"buyValueToCost"`
	SellValueToCostEr      float64 `json:"sellValueToCostEr"`
	SellValueToCost        float64 `json:"sellValueToCost"`
	MarkPriceEp            float64 `json:"markPriceEp"`
	MarkPrice              float64 `json:"markPrice"`
	MarkValueEv            float64 `json:"markValueEv"`
	MarkValue              float64 `json:"markValue"`
	UnRealisedPosLossEv    float64 `json:"unRealisedPosLossEv"`
	UnRealisedPosLoss      float64 `json:"unRealisedPosLoss"`
	EstimatedOrdLossEv     float64 `json:"estimatedOrdLossEv"`
	EstimatedOrdLoss       float64 `json:"estimatedOrdLoss"`
	UsedBalanceEv          float64 `json:"usedBalanceEv"`
	UsedBalance            float64 `json:"usedBalance"`
	TakeProfitEp           float64 `json:"takeProfitEp"`
	TakeProfit             float64 `json:"takeProfit"`
	StopLossEp             float64 `json:"stopLossEp"`
	StopLoss               float64 `json:"stopLoss"`
	RealisedPnlEv          float64 `json:"realisedPnlEv"`
	RealisedPnl            float64 `json:"realisedPnl"`
	CumRealisedPnlEv       float64 `json:"cumRealisedPnlEv"`
	CumRealisedPnl         float64 `json:"cumRealisedPnl"`
}
