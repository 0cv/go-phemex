package phemex

import (
	"context"
	"encoding/json"
)

// GetUserWalletService get account info
type GetUserWalletService struct {
	c         *Client
	offset    *int64
	limit     *int64
	withCount *int64
}

// Offset set offset
func (s *GetUserWalletService) Offset(offset int64) *GetUserWalletService {
	s.offset = &offset
	return s
}

// Limit set limit
func (s *GetUserWalletService) Limit(limit int64) *GetUserWalletService {
	s.limit = &limit
	return s
}

// WithCount set with count
func (s *GetUserWalletService) WithCount(withCount int64) *GetUserWalletService {
	s.withCount = &withCount
	return s
}

// Do send request
func (s *GetUserWalletService) Do(ctx context.Context, opts ...RequestOption) ([]*UserWallet, error) {
	r := &request{
		method:   "GET",
		endpoint: "/phemex-user/users/children",
		secType:  secTypeSigned,
	}
	if s.offset != nil {
		r.setParam("offset", s.offset)
	}
	if s.limit != nil {
		r.setParam("limit", s.limit)
	}
	if s.withCount != nil {
		r.setParam("withCount", s.withCount)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	resp := new(BaseResponse)
	resp.Data = new([]*UserWallet)
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	return *resp.Data.(*[]*UserWallet), nil
}

// UserWallet user wallet
type UserWallet struct {
	UserID        int                       `json:"userId"`
	Email         string                    `json:"email"`
	NickName      string                    `json:"nickName"`
	PasswordState float64                   `json:"passwordState"`
	ClientCnt     float64                   `json:"clientCnt"`
	Totp          float64                   `json:"totp"`
	Logon         float64                   `json:"logon"`
	ParentID      float64                   `json:"parentId"`
	ParentEmail   string                    `json:"parentEmail"`
	Status        float64                   `json:"status"`
	Wallet        *WalletUserWallet         `json:"wallet"`
	UserMarginVo  []*UserMarginUserWalletVo `json:"userMarginVo"`
}

// WalletUserWallet wallet user wallet
type WalletUserWallet struct {
	TotalBalance    string  `json:"totalBalance"`
	TotalBalanceEv  float64 `json:"totalBalanceEv"`
	AvailBalance    string  `json:"availBalance"`
	AvailBalanceEv  float64 `json:"availBalanceEv"`
	FreezeBalance   string  `json:"freezeBalance"`
	FreezeBalanceEv float64 `json:"freezeBalanceEv"`
	Currency        string  `json:"currency"`
	CurrencyCode    float64 `json:"currencyCode"`
}

// UserMarginUserWalletVo user margin UserWalletvo
type UserMarginUserWalletVo struct {
	Currency           string  `json:"currency"`
	AccountBalance     string  `json:"accountBalance"`
	TotalUsedBalance   string  `json:"totalUsedBalance"`
	AccountBalanceEv   float64 `json:"accountBalanceEv"`
	TotalUsedBalanceEv float64 `json:"totalUsedBalanceEv"`
	BonusBalanceEv     float64 `json:"bonusBalanceEv"`
	BonusBalance       string  `json:"bonusBalance"`
}

// ExchangeMarginService get account info
type ExchangeMarginService struct {
	c           *Client
	btcAmount   *float64
	btcAmountEv *int64
	linkKey     *string
	moveOp      ExchangeMarginType
	usdAmount   *float64
	usdAmountEv *int64
}

// BtcAmount set btcAmount
func (s *ExchangeMarginService) BtcAmount(btcAmount float64) *ExchangeMarginService {
	s.btcAmount = &btcAmount
	return s
}

// BtcAmountEv set btcAmountEv
func (s *ExchangeMarginService) BtcAmountEv(btcAmountEv int64) *ExchangeMarginService {
	s.btcAmountEv = &btcAmountEv
	return s
}

// LinkKey set linkKey
func (s *ExchangeMarginService) LinkKey(linkKey string) *ExchangeMarginService {
	s.linkKey = &linkKey
	return s
}

// MoveOp set moveOp
func (s *ExchangeMarginService) MoveOp(moveOp ExchangeMarginType) *ExchangeMarginService {
	s.moveOp = moveOp
	return s
}

// UsdAmount set usdAmount
func (s *ExchangeMarginService) UsdAmount(usdAmount float64) *ExchangeMarginService {
	s.usdAmount = &usdAmount
	return s
}

// UsdAmountEv set usdAmountEv
func (s *ExchangeMarginService) UsdAmountEv(usdAmountEv int64) *ExchangeMarginService {
	s.usdAmountEv = &usdAmountEv
	return s
}

// Do send request
func (s *ExchangeMarginService) Do(ctx context.Context, opts ...RequestOption) (*ExchangeMargin, error) {
	r := &request{
		method:   "POST",
		endpoint: "/exchange/margins",
		secType:  secTypeSigned,
	}

	m := params{
		"moveOp": s.moveOp,
	}

	if s.btcAmount != nil {
		m["btcAmount"] = *s.btcAmount
	}
	if s.btcAmountEv != nil {
		m["btcAmountEv"] = *s.btcAmountEv
	}
	if s.linkKey != nil {
		m["linkKey"] = *s.linkKey
	}
	if s.usdAmount != nil {
		m["usdAmount"] = *s.usdAmount
	}
	if s.usdAmountEv != nil {
		m["usdAmountEv"] = *s.usdAmountEv
	}
	r.setFormParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	resp := new(BaseResponse)
	resp.Data = new(ExchangeMargin)
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	return resp.Data.(*ExchangeMargin), nil
}

// ExchangeMargin exchange margin
type ExchangeMargin struct {
	MoveOp           int    `json:"moveOp"`
	FromCurrencyName string `json:"fromCurrencyName"`
	ToCurrencyName   string `json:"toCurrencyName"`
	FromAmount       string `json:"fromAmount"`
	ToAmount         string `json:"toAmount"`
	LinkKey          string `json:"linkKey"`
	Status           int64  `json:"status"`
}
