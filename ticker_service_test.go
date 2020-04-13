package phemex

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type tickerServiceTestSuite struct {
	baseTestSuite
}

func TestTickerService(t *testing.T) {
	suite.Run(t, new(tickerServiceTestSuite))
}

func (s *tickerServiceTestSuite) TestTickerService() {
	data := []byte(`{
		"error": null,
		"id": 0,
		"result": [{
			"askEp": 0,
			"bidEp": 0,
			"openEp": 0,
			"highEp": 0,
			"lowEp": 0,
			"lastEp": 0,
			"indexEp": 0,
			"markEp": 0,
			"openInterest": 0,
			"fundingRateEr": 0,
			"predFundingRateEr": 0,
			"timestamp": 11111,
			"symbol": "BTCUSD",
			"volume": 0
		}]
	}`)

	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol": "BTCUSD",
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewListPriceChangeStatsService().Symbol("BTCUSD").Do(newContext())
	s.r().NoError(err)
	e := []*PriceChangeStats{{
		AskEp:             0,
		BidEp:             0,
		OpenEp:            0,
		HighEp:            0,
		LowEp:             0,
		LastEp:            0,
		IndexEp:           0,
		MarkEp:            0,
		OpenInterest:      0,
		FundingRateEr:     0,
		PredFundingRateEr: 0,
		Timestamp:         11111,
		Symbol:            "BTCUSD",
		Volume:            0,
	}}
	s.assertAccountEqual(e, res)
}

func (s *tickerServiceTestSuite) assertAccountEqual(e, a []*PriceChangeStats) {
	r := s.r()
	for i := 0; i < len(a); i++ {
		r.Equal(e[i].AskEp, a[i].AskEp, "AskEp")
		r.Equal(e[i].BidEp, a[i].BidEp, "BidEp")
		r.Equal(e[i].OpenEp, a[i].OpenEp, "OpenEp")
		r.Equal(e[i].HighEp, a[i].HighEp, "HighEp")
		r.Equal(e[i].LowEp, a[i].LowEp, "LowEp")
		r.Equal(e[i].LastEp, a[i].LastEp, "LastEp")
		r.Equal(e[i].IndexEp, a[i].IndexEp, "IndexEp")
		r.Equal(e[i].MarkEp, a[i].MarkEp, "MarkEp")
		r.Equal(e[i].OpenInterest, a[i].OpenInterest, "OpenInterest")
		r.Equal(e[i].FundingRateEr, a[i].FundingRateEr, "FundingRateEr")
		r.Equal(e[i].PredFundingRateEr, a[i].PredFundingRateEr, "PredFundingRateEr")
		r.Equal(e[i].Timestamp, a[i].Timestamp, "Timestamp")
		r.Equal(e[i].Symbol, a[i].Symbol, "Symbol")
		r.Equal(e[i].Volume, a[i].Volume, "Volume")
	}
}
