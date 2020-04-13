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
func (s *ExchangeProductsService) Do(ctx context.Context, opts ...RequestOption) (res []*ExchangeProduct, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/exchange/public/products",
		secType:  secTypeNone,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	//res = new(ExchangeProduct)
	resp := new(BaseResponse)
	resp.Data = new([]*ExchangeProduct)
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}

	if resp.Data == nil {
		return nil, errors.New("Null response")
	}

	rows := resp.Data.(*[]*ExchangeProduct)
	return *rows, nil
}

// ExchangeProduct exchange products
type ExchangeProduct struct {
	Symbol             string  `json:"symbol"`
	UnderlyingSymbol   string  `json:"underlyingSymbol"`
	QuoteCurrency      string  `json:"quoteCurrency"`
	SettlementCurrency string  `json:"settlementCurrency"`
	MaxOrderQty        float64 `json:"maxOrderQty"`
	LotSize            float64 `json:"lotSize"`
	TickSize           string  `json:"tickSize"`
	ContractSize       string  `json:"contractSize"`
	PriceScale         float64 `json:"priceScale"`
	RatioScale         float64 `json:"ratioScale"`
	ValueScale         float64 `json:"valueScale"`
	DefaultLeverage    float64 `json:"defaultLeverage"`
	MaxLeverage        float64 `json:"maxLeverage"`
	InitMarginEr       string  `json:"initMarginEr"`
	MaintMarginEr      string  `json:"maintMarginEr"`
	DefaultRiskLimitEv float64 `json:"defaultRiskLimitEv"`
	Deleverage         bool    `json:"deleverage"`
	MakerFeeRateEr     int64   `json:"makerFeeRateEr"`
	TakerFeeRateEr     int64   `json:"takerFeeRateEr"`
	FundingInterval    float64 `json:"fundingInterval"`
	Description        string  `json:"description"`
}
