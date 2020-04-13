package phemex

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type exchangeInfoServiceTestSuite struct {
	baseTestSuite
}

func TestExchangeInfoService(t *testing.T) {
	suite.Run(t, new(exchangeInfoServiceTestSuite))
}

func (s *exchangeInfoServiceTestSuite) TestExchangeInfoService() {
	data := []byte(`{
		"code": 0,
		"msg": "OK",
		"data": [
		  {
			"symbol": "BTC",
			"underlyingSymbol": "BTC",
			"quoteCurrency": "BTC",
			"settlementCurrency": "BTC",
			"maxOrderQty": 0,
			"maxPriceEp": 0,
			"lotSize": 0,
			"tickSize": "1",
			"contractSize": "1",
			"priceScale": 0,
			"ratioScale": 0,
			"valueScale": 0,
			"defaultLeverage": 0,
			"maxLeverage": 0,
			"initMarginEr": "1",
			"maintMarginEr": "1",
			"defaultRiskLimitEv": 0,
			"deleverage": false,
			"makerFeeRateEr": 0,
			"takerFeeRateEr": 0,
			"fundingInterval": 0,
			"description": ""
		  }
		]
	}`)

	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewExchangeProductsService().Do(newContext())
	s.r().NoError(err)
	e := []*ExchangeProduct{{
		Symbol:             "BTC",
		UnderlyingSymbol:   "BTC",
		QuoteCurrency:      "BTC",
		SettlementCurrency: "BTC",
		MaxOrderQty:        0,
		LotSize:            0,
		TickSize:           "1",
		ContractSize:       "1",
		PriceScale:         0,
		RatioScale:         0,
		ValueScale:         0,
		DefaultLeverage:    0,
		MaxLeverage:        0,
		InitMarginEr:       "1",
		MaintMarginEr:      "1",
		DefaultRiskLimitEv: 0,
		Deleverage:         false,
		MakerFeeRateEr:     0,
		TakerFeeRateEr:     0,
		FundingInterval:    0,
		Description:        "",
	}}
	s.assertAccountEqual(e, res)
}

func (s *exchangeInfoServiceTestSuite) assertAccountEqual(e, a []*ExchangeProduct) {
	r := s.r()
	for i := 0; i < len(a); i++ {
		r.Equal(e[i].Symbol, a[i].Symbol, "Symbol")
		r.Equal(e[i].UnderlyingSymbol, a[i].UnderlyingSymbol, "UnderlyingSymbol")
		r.Equal(e[i].QuoteCurrency, a[i].QuoteCurrency, "QuoteCurrency")
		r.Equal(e[i].SettlementCurrency, a[i].SettlementCurrency, "SettlementCurrency")
		r.Equal(e[i].MaxOrderQty, a[i].MaxOrderQty, "MaxOrderQty")
		r.Equal(e[i].LotSize, a[i].LotSize, "LotSize")
		r.Equal(e[i].TickSize, a[i].TickSize, "TickSize")
		r.Equal(e[i].ContractSize, a[i].ContractSize, "ContractSize")
		r.Equal(e[i].PriceScale, a[i].PriceScale, "PriceScale")
		r.Equal(e[i].RatioScale, a[i].RatioScale, "RatioScale")
		r.Equal(e[i].ValueScale, a[i].ValueScale, "ValueScale")
		r.Equal(e[i].DefaultLeverage, a[i].DefaultLeverage, "DefaultLeverage")
		r.Equal(e[i].MaxLeverage, a[i].MaxLeverage, "MaxLeverage")
		r.Equal(e[i].InitMarginEr, a[i].InitMarginEr, "InitMarginEr")
		r.Equal(e[i].MaintMarginEr, a[i].MaintMarginEr, "MaintMarginEr")
		r.Equal(e[i].DefaultRiskLimitEv, a[i].DefaultRiskLimitEv, "DefaultRiskLimitEv")
		r.Equal(e[i].Deleverage, a[i].Deleverage, "Deleverage")
		r.Equal(e[i].MakerFeeRateEr, a[i].MakerFeeRateEr, "MakerFeeRateEr")
		r.Equal(e[i].TakerFeeRateEr, a[i].TakerFeeRateEr, "TakerFeeRateEr")
		r.Equal(e[i].FundingInterval, a[i].FundingInterval, "FundingInterval")
		r.Equal(e[i].Description, a[i].Description, "Description")
	}
}
