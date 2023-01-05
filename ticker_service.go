package phemex

import (
	"context"
	"encoding/json"
)

// ListPriceChangeStatsService show stats of price change in last 24 hours for all symbols
type ListPriceChangeStatsService struct {
	c      *Client
	symbol *string
}

// Symbol set symbol
func (s *ListPriceChangeStatsService) Symbol(symbol string) *ListPriceChangeStatsService {
	s.symbol = &symbol
	return s
}

// Do send request
func (s *ListPriceChangeStatsService) Do(ctx context.Context, opts ...RequestOption) (res *PriceChangeStats, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/v1/md/ticker/24hr/all",
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}
	resp := new(BaseTickerResponse)
	resp.Result = new(PriceChangeStats)

	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}

	return resp.Result.(*PriceChangeStats), nil
}

// BaseTickerResponse base ticker response
type BaseTickerResponse struct {
	Error  int64       `json:"error"`
	ID     int64       `json:"id"`
	Result interface{} `json:"result"`
}

// PriceChangeStats define price change stats
type PriceChangeStats struct {
	AskEp             int64   `json:"askEp"`
	BidEp             int64   `json:"bidEp"`
	OpenEp            int64   `json:"openEp"`
	HighEp            int64   `json:"highEp"`
	LowEp             int64   `json:"lowEp"`
	LastEp            int64   `json:"lastEp"`
	IndexEp           int64   `json:"indexEp"`
	MarkEp            int64   `json:"markEp"`
	OpenInterest      float64 `json:"openInterest"`
	FundingRateEr     float64 `json:"fundingRateEr"`
	PredFundingRateEr float64 `json:"predFundingRateEr"`
	Timestamp         int64   `json:"timestamp"`
	Symbol            string  `json:"symbol"`
	TurnoverEv        int64   `json:"turnoverEv"`
	Volume            float64 `json:"volume"`
}
