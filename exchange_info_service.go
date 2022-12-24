package phemex

import (
	"context"
	"encoding/json"
	"errors"
)

// ExchangeProductsService exchange info service
type ExchangeProductsService struct {
	c *Client
}

// Do send request
func (s *ExchangeProductsService) Do(ctx context.Context, opts ...RequestOption) (res *ExchangeProductsServiceResponse, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/exchange/public/cfg/v2/products",
		secType:  secTypeNone,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	resp := new(BaseResponse)
	resp.Data = &ExchangeProductsServiceResponse{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}

	if resp.Data == nil {
		return nil, errors.New("Null response")
	}

	return resp.Data.(*ExchangeProductsServiceResponse), nil
}

type ExchangeProductsServiceResponse struct {
	Currencies []ExchangeCurrency  `json:"currencies"`
	Products   []ExchangeProduct   `json:"products"`
	RiskLimits []ExchangeRiskLimit `json:"riskLimitsV2"`
	Leverages  []ExchangeLeverage  `json:"leverages"`
}

type ExchangeCurrency struct {
}

type ExchangeRiskLimit struct{}

type ExchangeLeverage struct{}

type ExchangeProduct struct {
	Symbol                   string  `json:"symbol"`
	Code                     int64   `json:"code"`
	DisplaySymbol            string  `json:"displaySymbol"`
	IndexSymbol              string  `json:"indexSymbol"`
	MarkSymbol               string  `json:"markSymbol"`
	FundingRateSymbol        string  `json:"fundingRateSymbol"`
	FundingRate8HSymbol      string  `json:"fundingRate8hSymbol"`
	ContractUnderlyingAssets string  `json:"contractUnderlyingAssets"`
	SettleCurrency           string  `json:"settleCurrency"`
	QuoteCurrency            string  `json:"quoteCurrency"`
	ContractSize             float64 `json:"contractSize"`
	LotSize                  int64   `json:"lotSize"`
	TickSize                 float64 `json:"tickSize"`
	PriceScale               int64   `json:"priceScale"`
	RatioScale               int64   `json:"ratioScale"`
	PricePrecision           int64   `json:"pricePrecision"`
	MinPriceEp               int64   `json:"minPriceEp"`
	MaxPriceEp               int64   `json:"maxPriceEp"`
	MaxOrderQty              int64   `json:"maxOrderQty"`
	Type                     string  `json:"type"`
	Status                   string  `json:"status"`
	TipOrderQty              int64   `json:"tipOrderQty"`
	Description              string  `json:"description"`
	//PerpetualV2
	//QtyPrecision    int     `json:"qtyPrecision"`
	//QtyStepSize     float64 `json:"qtyStepSize"`
	//MinPriceRp      float64 `json:"minPriceRp"`
	//MaxPriceRp      float64 `json:"maxPriceRp"`
	//MinOrderValueRv float64 `json:"minOrderValueRv"`
	//MaxOrderQtyRq   float64 `json:"maxOrderQtyRq"`
	//BaseCurrency    string  `json:"baseCurrency"`
	//TipOrderQtyRq   float64 `json:"tipOrderQtyRq"`
}
