package phemex

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type walletServiceTestSuite struct {
	baseTestSuite
}

func TestWalletService(t *testing.T) {
	suite.Run(t, new(walletServiceTestSuite))
}

func (s *walletServiceTestSuite) TestUserWalletService() {
	data := []byte(`{
		"code": 0,
		"msg": "OK",
		"data": [{
			"userId": 123123,
			"email": "x**@**.com",
			"nickName": "nickName",
			"passwordState": 1,
			"clientCnt": 0,
			"totp": 1,
			"logon": 0,
			"parentId": 0,
			"parentEmail": null,
			"status": 1,
			"wallet": {
				"totalBalance": "989.25471319",
				"totalBalanceEv": 98925471319,
				"availBalance": "989.05471319",
				"availBalanceEv": 98905471319,
				"freezeBalance": "0.20000000",
				"freezeBalanceEv": 20000000,
				"currency": "BTC",
				"currencyCode": 1
			},
			"userMarginVo": [{
				"currency": "BTC",
				"accountBalance": "3.90032508",
				"totalUsedBalance": "0.00015666",
				"accountBalanceEv": 390032508,
				"totalUsedBalanceEv": 15666,
				"bonusBalanceEv": 0,
				"bonusBalance": "0"
			},{
				"currency": "USD",
				"accountBalance": "38050.35000000",
				"totalUsedBalance": "0.00000000",
				"accountBalanceEv": 380503500,
				"totalUsedBalanceEv": 0,
				"bonusBalanceEv": 0,
				"bonusBalance": "0"
			}]
		}]
	}`)

	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetUserWalletService().Do(newContext())
	s.r().NoError(err)
	e := []*UserWallet{{
		UserID:        123123,
		Email:         "x**@**.com",
		NickName:      "nickName",
		PasswordState: 1,
		ClientCnt:     0,
		Totp:          1,
		Logon:         0,
		ParentID:      0,
		ParentEmail:   "",
		Status:        1,
		Wallet: &WalletUserWallet{
			TotalBalance:    "989.25471319",
			TotalBalanceEv:  98925471319,
			AvailBalance:    "989.05471319",
			AvailBalanceEv:  98905471319,
			FreezeBalance:   "0.20000000",
			FreezeBalanceEv: 20000000,
			Currency:        "BTC",
			CurrencyCode:    1,
		},
		UserMarginVo: []*UserMarginUserWalletVo{{
			Currency:           "BTC",
			AccountBalance:     "3.90032508",
			TotalUsedBalance:   "0.00015666",
			AccountBalanceEv:   390032508,
			TotalUsedBalanceEv: 15666,
			BonusBalanceEv:     0,
			BonusBalance:       "0",
		}},
	}}
	s.assertAccountEqual(e, res)
}

func (s *walletServiceTestSuite) assertAccountEqual(e, a []*UserWallet) {
	r := s.r()
	for i := 0; i < len(a); i++ {
		r.Equal(e[i].UserID, a[i].UserID, "UserID")
		r.Equal(e[i].Email, a[i].Email, "Email")
		r.Equal(e[i].NickName, a[i].NickName, "NickName")
		r.Equal(e[i].PasswordState, a[i].PasswordState, "PasswordState")
		r.Equal(e[i].ClientCnt, a[i].ClientCnt, "ClientCnt")
		r.Equal(e[i].Totp, a[i].Totp, "Totp")
		r.Equal(e[i].Logon, a[i].Logon, "Logon")
		r.Equal(e[i].ParentID, a[i].ParentID, "ParentID")
		r.Equal(e[i].ParentEmail, a[i].ParentEmail, "ParentEmail")
		r.Equal(e[i].Status, a[i].Status, "Status")
		r.Equal(e[i].Wallet.TotalBalance, a[i].Wallet.TotalBalance, "TotalBalance")
		r.Equal(e[i].Wallet.TotalBalanceEv, a[i].Wallet.TotalBalanceEv, "TotalBalanceEv")
		r.Equal(e[i].Wallet.AvailBalance, a[i].Wallet.AvailBalance, "AvailBalance")
		r.Equal(e[i].Wallet.AvailBalanceEv, a[i].Wallet.AvailBalanceEv, "AvailBalanceEv")
		r.Equal(e[i].Wallet.FreezeBalance, a[i].Wallet.FreezeBalance, "FreezeBalance")
		r.Equal(e[i].Wallet.FreezeBalanceEv, a[i].Wallet.FreezeBalanceEv, "FreezeBalanceEv")
		r.Equal(e[i].Wallet.Currency, a[i].Wallet.Currency, "Currency")
		r.Equal(e[i].Wallet.CurrencyCode, a[i].Wallet.CurrencyCode, "CurrencyCode")
		for j := 0; j < len(a); j++ {
			r.Equal(e[i].UserMarginVo[j].Currency, a[i].UserMarginVo[j].Currency, "Currency")
			r.Equal(e[i].UserMarginVo[j].AccountBalance, a[i].UserMarginVo[j].AccountBalance, "AccountBalance")
			r.Equal(e[i].UserMarginVo[j].TotalUsedBalance, a[i].UserMarginVo[j].TotalUsedBalance, "TotalUsedBalance")
			r.Equal(e[i].UserMarginVo[j].AccountBalanceEv, a[i].UserMarginVo[j].AccountBalanceEv, "AccountBalanceEv")
			r.Equal(e[i].UserMarginVo[j].TotalUsedBalanceEv, a[i].UserMarginVo[j].TotalUsedBalanceEv, "TotalUsedBalanceEv")
			r.Equal(e[i].UserMarginVo[j].BonusBalanceEv, a[i].UserMarginVo[j].BonusBalanceEv, "BonusBalanceEv")
			r.Equal(e[i].UserMarginVo[j].BonusBalance, a[i].UserMarginVo[j].BonusBalance, "BonusBalance")
		}
	}
}

func (s *walletServiceTestSuite) TestExchangeMarginService() {
	data := []byte(`{
		"code": 0,
		"msg": "OK",
		"data": {
			"moveOp": 0,
			"fromCurrencyName": "BTC",
			"toCurrencyName": "BTC",
			"fromAmount": "0.10000000",
			"toAmount": "0.10000000",
			"linkKey": "2431ca9b-2dd4-44b8-91f3-2539bb62db2d",
			"status": 10
		}
	}`)

	s.mockDo(data, nil)
	defer s.assertDo()
	moveOp := float64(0)
	btcAmount := float64(10)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"moveOp":    moveOp,
			"btcAmount": btcAmount,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewExchangeMarginService().MoveOp(0).BtcAmount(btcAmount).Do(newContext())
	s.r().NoError(err)
	e := &ExchangeMargin{
		MoveOp:           0,
		FromCurrencyName: "BTC",
		ToCurrencyName:   "BTC",
		FromAmount:       "0.10000000",
		ToAmount:         "0.10000000",
		LinkKey:          "2431ca9b-2dd4-44b8-91f3-2539bb62db2d",
		Status:           10,
	}
	s.assertExchangeMarginEquals(e, res)
}

func (s *walletServiceTestSuite) assertExchangeMarginEquals(e, a *ExchangeMargin) {
	r := s.r()
	r.Equal(e.MoveOp, a.MoveOp, "MoveOp")
	r.Equal(e.FromCurrencyName, a.FromCurrencyName, "FromCurrencyName")
	r.Equal(e.ToCurrencyName, a.ToCurrencyName, "ToCurrencyName")
	r.Equal(e.FromAmount, a.FromAmount, "FromAmount")
	r.Equal(e.ToAmount, a.ToAmount, "ToAmount")
	r.Equal(e.LinkKey, a.LinkKey, "LinkKey")
	r.Equal(e.Status, a.Status, "Status")
}
