package phemex

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type accountServiceTestSuite struct {
	baseTestSuite
}

func TestAccountService(t *testing.T) {
	suite.Run(t, new(accountServiceTestSuite))
}

func (s *accountServiceTestSuite) TestGetAccount() {
	data := []byte(`{
		"code": 0,
		"msg": "",
		"data": {
			"account": {
				"accountId": 0,
				"currency": "BTC",
				"accountBalanceEv": 0,
				"totalUsedBalanceEv": 0
			},
			"positions": [
			{
				"accountID": 0,
				"symbol": "BTCUSD",
				"currency": "BTC",
				"side": "None",
				"positionStatus": "Normal",
				"crossMargin": false,
				"leverageEr": 0,
				"leverage": 0,
				"initMarginReqEr": 0,
				"initMarginReq": 0.01,
				"maintMarginReqEr": 500000,
				"maintMarginReq": 0.005,
				"riskLimitEv": 10000000000,
				"riskLimit": 100,
				"size": 0,
				"value": 0,
				"valueEv": 0,
				"avgEntryPriceEp": 0,
				"avgEntryPrice": 0,
				"posCostEv": 0,
				"posCost": 0,
				"assignedPosBalanceEv": 0,
				"assignedPosBalance": 0,
				"bankruptCommEv": 0,
				"bankruptComm": 0,
				"bankruptPriceEp": 0,
				"bankruptPrice": 0,
				"positionMarginEv": 0,
				"positionMargin": 0,
				"liquidationPriceEp": 0,
				"liquidationPrice": 0,
				"deleveragePercentileEr": 0,
				"deleveragePercentile": 0,
				"buyValueToCostEr": 1150750,
				"buyValueToCost": 0.0115075,
				"sellValueToCostEr": 1149250,
				"sellValueToCost": 0.0114925,
				"markPriceEp": 93169002,
				"markPrice": 9316.9002,
				"markValueEv": 0,
				"markValue": 0,
				"unRealisedPosLossEv": 0,
				"unRealisedPosLoss": 0,
				"estimatedOrdLossEv": 0,
				"estimatedOrdLoss": 0,
				"usedBalanceEv": 0,
				"usedBalance": 0,
				"takeProfitEp": 0,
				"takeProfit": 0,
				"stopLossEp": 0,
				"stopLoss": 0,
				"realisedPnlEv": 0,
				"realisedPnl": 0,
				"cumRealisedPnlEv": 0,
				"cumRealisedPnl": 0
			}
			]
		}
	}`)

	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetAccountPositionService().Currency("BTC").Do(newContext())
	s.r().NoError(err)
	e := &AccountPosition{
		Account: Account{
			AccountID:          0,
			Currency:           "BTC",
			AccountBalanceEv:   0,
			TotalUsedBalanceEv: 0,
		},
		Positions: []*Position{{
			AccountID:              0,
			Symbol:                 "BTCUSD",
			Currency:               "BTC",
			Side:                   "None",
			PositionStatus:         "Normal",
			CrossMargin:            false,
			LeverageEr:             0,
			Leverage:               0,
			InitMarginReqEr:        0,
			InitMarginReq:          0.01,
			MaintMarginReqEr:       500000,
			MaintMarginReq:         0.005,
			RiskLimitEv:            10000000000,
			RiskLimit:              100,
			Size:                   0,
			Value:                  0,
			ValueEv:                0,
			AvgEntryPriceEp:        0,
			AvgEntryPrice:          0,
			PosCostEv:              0,
			PosCost:                0,
			AssignedPosBalanceEv:   0,
			AssignedPosBalance:     0,
			BankruptCommEv:         0,
			BankruptComm:           0,
			BankruptPriceEp:        0,
			BankruptPrice:          0,
			PositionMarginEv:       0,
			PositionMargin:         0,
			LiquidationPriceEp:     0,
			LiquidationPrice:       0,
			DeleveragePercentileEr: 0,
			DeleveragePercentile:   0,
			BuyValueToCostEr:       1150750,
			BuyValueToCost:         0.0115075,
			SellValueToCostEr:      1149250,
			SellValueToCost:        0.0114925,
			MarkPriceEp:            93169002,
			MarkPrice:              9316.9002,
			MarkValueEv:            0,
			MarkValue:              0,
			UnRealisedPosLossEv:    0,
			UnRealisedPosLoss:      0,
			EstimatedOrdLossEv:     0,
			EstimatedOrdLoss:       0,
			UsedBalanceEv:          0,
			UsedBalance:            0,
			TakeProfitEp:           0,
			TakeProfit:             0,
			StopLossEp:             0,
			StopLoss:               0,
			RealisedPnlEv:          0,
			RealisedPnl:            0,
			CumRealisedPnlEv:       0,
			CumRealisedPnl:         0,
		}},
	}
	s.assertAccountEqual(e, res)
}

func (s *accountServiceTestSuite) assertAccountEqual(e, a *AccountPosition) {
	r := s.r()
	r.Equal(e.Account.AccountID, a.Account.AccountID, "AccountId")
	r.Equal(e.Account.Currency, a.Account.Currency, "Currency")
	r.Equal(e.Account.AccountBalanceEv, a.Account.AccountBalanceEv, "AccountBalanceEv")
	r.Equal(e.Account.TotalUsedBalanceEv, a.Account.TotalUsedBalanceEv, "TotalUsedBalanceEv")
	for i := 0; i < len(a.Positions); i++ {
		r.Equal(e.Positions[i].AccountID, a.Positions[i].AccountID, "AccountID")
		r.Equal(e.Positions[i].Symbol, a.Positions[i].Symbol, "Symbol")
		r.Equal(e.Positions[i].Currency, a.Positions[i].Currency, "Currency")
		r.Equal(e.Positions[i].Side, a.Positions[i].Side, "Side")
		r.Equal(e.Positions[i].PositionStatus, a.Positions[i].PositionStatus, "PositionStatus")
		r.Equal(e.Positions[i].CrossMargin, a.Positions[i].CrossMargin, "CrossMargin")
		r.Equal(e.Positions[i].LeverageEr, a.Positions[i].LeverageEr, "LeverageEr")
		r.Equal(e.Positions[i].Leverage, a.Positions[i].Leverage, "Leverage")
		r.Equal(e.Positions[i].InitMarginReqEr, a.Positions[i].InitMarginReqEr, "InitMarginReqEr")
		r.Equal(e.Positions[i].InitMarginReq, a.Positions[i].InitMarginReq, "InitMarginReq")
		r.Equal(e.Positions[i].MaintMarginReqEr, a.Positions[i].MaintMarginReqEr, "MaintMarginReqEr")
		r.Equal(e.Positions[i].MaintMarginReq, a.Positions[i].MaintMarginReq, "MaintMarginReq")
		r.Equal(e.Positions[i].RiskLimitEv, a.Positions[i].RiskLimitEv, "RiskLimitEv")
		r.Equal(e.Positions[i].RiskLimit, a.Positions[i].RiskLimit, "RiskLimit")
		r.Equal(e.Positions[i].Size, a.Positions[i].Size, "Size")
		r.Equal(e.Positions[i].Value, a.Positions[i].Value, "Value")
		r.Equal(e.Positions[i].ValueEv, a.Positions[i].ValueEv, "ValueEv")
		r.Equal(e.Positions[i].AvgEntryPriceEp, a.Positions[i].AvgEntryPriceEp, "AvgEntryPriceEp")
		r.Equal(e.Positions[i].AvgEntryPrice, a.Positions[i].AvgEntryPrice, "AvgEntryPrice")
		r.Equal(e.Positions[i].PosCostEv, a.Positions[i].PosCostEv, "PosCostEv")
		r.Equal(e.Positions[i].PosCost, a.Positions[i].PosCost, "PosCost")
		r.Equal(e.Positions[i].AssignedPosBalanceEv, a.Positions[i].AssignedPosBalanceEv, "AssignedPosBalanceEv")
		r.Equal(e.Positions[i].AssignedPosBalance, a.Positions[i].AssignedPosBalance, "AssignedPosBalance")
		r.Equal(e.Positions[i].BankruptCommEv, a.Positions[i].BankruptCommEv, "BankruptCommEv")
		r.Equal(e.Positions[i].BankruptComm, a.Positions[i].BankruptComm, "BankruptComm")
		r.Equal(e.Positions[i].BankruptPriceEp, a.Positions[i].BankruptPriceEp, "BankruptPriceEp")
		r.Equal(e.Positions[i].BankruptPrice, a.Positions[i].BankruptPrice, "BankruptPrice")
		r.Equal(e.Positions[i].PositionMarginEv, a.Positions[i].PositionMarginEv, "PositionMarginEv")
		r.Equal(e.Positions[i].PositionMargin, a.Positions[i].PositionMargin, "PositionMargin")
		r.Equal(e.Positions[i].LiquidationPriceEp, a.Positions[i].LiquidationPriceEp, "LiquidationPriceEp")
		r.Equal(e.Positions[i].LiquidationPrice, a.Positions[i].LiquidationPrice, "LiquidationPrice")
		r.Equal(e.Positions[i].DeleveragePercentileEr, a.Positions[i].DeleveragePercentileEr, "DeleveragePercentileEr")
		r.Equal(e.Positions[i].DeleveragePercentile, a.Positions[i].DeleveragePercentile, "DeleveragePercentile")
		r.Equal(e.Positions[i].BuyValueToCostEr, a.Positions[i].BuyValueToCostEr, "BuyValueToCostEr")
		r.Equal(e.Positions[i].BuyValueToCost, a.Positions[i].BuyValueToCost, "BuyValueToCost")
		r.Equal(e.Positions[i].SellValueToCostEr, a.Positions[i].SellValueToCostEr, "SellValueToCostEr")
		r.Equal(e.Positions[i].SellValueToCost, a.Positions[i].SellValueToCost, "SellValueToCost")
		r.Equal(e.Positions[i].MarkPriceEp, a.Positions[i].MarkPriceEp, "MarkPriceEp")
		r.Equal(e.Positions[i].MarkPrice, a.Positions[i].MarkPrice, "MarkPrice")
		r.Equal(e.Positions[i].MarkValueEv, a.Positions[i].MarkValueEv, "MarkValueEv")
		r.Equal(e.Positions[i].MarkValue, a.Positions[i].MarkValue, "MarkValue")
		r.Equal(e.Positions[i].UnRealisedPosLossEv, a.Positions[i].UnRealisedPosLossEv, "UnRealisedPosLossEv")
		r.Equal(e.Positions[i].UnRealisedPosLoss, a.Positions[i].UnRealisedPosLoss, "UnRealisedPosLoss")
		r.Equal(e.Positions[i].EstimatedOrdLossEv, a.Positions[i].EstimatedOrdLossEv, "EstimatedOrdLossEv")
		r.Equal(e.Positions[i].EstimatedOrdLoss, a.Positions[i].EstimatedOrdLoss, "EstimatedOrdLoss")
		r.Equal(e.Positions[i].UsedBalanceEv, a.Positions[i].UsedBalanceEv, "UsedBalanceEv")
		r.Equal(e.Positions[i].UsedBalance, a.Positions[i].UsedBalance, "UsedBalance")
		r.Equal(e.Positions[i].TakeProfitEp, a.Positions[i].TakeProfitEp, "TakeProfitEp")
		r.Equal(e.Positions[i].TakeProfit, a.Positions[i].TakeProfit, "TakeProfit")
		r.Equal(e.Positions[i].StopLossEp, a.Positions[i].StopLossEp, "StopLossEp")
		r.Equal(e.Positions[i].StopLoss, a.Positions[i].StopLoss, "StopLoss")
		r.Equal(e.Positions[i].RealisedPnlEv, a.Positions[i].RealisedPnlEv, "RealisedPnlEv")
		r.Equal(e.Positions[i].RealisedPnl, a.Positions[i].RealisedPnl, "RealisedPnl")
		r.Equal(e.Positions[i].CumRealisedPnl, a.Positions[i].CumRealisedPnl, "CumRealisedPnl")
		r.Equal(e.Positions[i].CumRealisedPnlEv, a.Positions[i].CumRealisedPnlEv, "CumRealisedPnlEv")
	}
}
