package phemex

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Krisa/go-phemex/common"
	"github.com/bitly/go-simplejson"
)

// ExchangeMarginType define exchange margin type
type ExchangeMarginType int

// SideType define side type of order
type SideType string

// OrderType define order type
type OrderType string

// TriggerType define trigger type
type TriggerType string

// TimeInForceType define time in force type of order
type TimeInForceType string

// Global enums
const (
	ExchangeMarginTypeBTCToWallet ExchangeMarginType = 1
	ExchangeMarginTypeWalletToBTC ExchangeMarginType = 2
	ExchangeMarginTypeWalletToUSD ExchangeMarginType = 3
	ExchangeMarginTypeUSDToWallet ExchangeMarginType = 4

	SideTypeBuy  SideType = "Buy"
	SideTypeSell SideType = "Sell"

	OrderTypeLimit           OrderType = "Limit"
	OrderTypeMarket          OrderType = "Market"
	OrderTypeStop            OrderType = "Stop"
	OrderTypeStopLimit       OrderType = "StopLimit"
	OrderTypeMarketIfTouched OrderType = "MarketIfTouched"
	OrderTypeLimitIfTouched  OrderType = "LimitIfTouched"
	OrderTypePegged          OrderType = "Pegged"

	TimeInForceTypeDAY TimeInForceType = "Day"
	TimeInForceTypeGTC TimeInForceType = "GoodTillCancel"
	TimeInForceTypeIOC TimeInForceType = "ImmediateOrCancel"
	TimeInForceTypeFOK TimeInForceType = "FillOrKill"

	TriggerTypeByMarkPrice TriggerType = "ByMarkPrice"
	TriggerTypeByLastPrice TriggerType = "ByLastPrice"

	signatureKey   = "x-phemex-request-signature"
	expiryKey      = "x-phemex-request-expiry"
	accessTokenKey = "x-phemex-access-token"
)

func currentTimestamp() int64 {
	return time.Now().Unix()
}

func newJSON(data []byte) (j *simplejson.Json, err error) {
	j, err = simplejson.NewJson(data)
	if err != nil {
		return nil, err
	}
	return j, nil
}

// NewClient initialize an API client instance with API key and secret key.
// You should always call this function before using this SDK.
// Services will be created by the form client.NewXXXService().
func NewClient(apiKey, secretKey string) *Client {
	return &Client{
		APIKey:     apiKey,
		SecretKey:  secretKey,
		BaseURL:    "https://api.phemex.com",
		UserAgent:  "Phemex/golang",
		HTTPClient: http.DefaultClient,
		Logger:     log.New(os.Stderr, "Phemex-golang", log.LstdFlags),
	}
}

type doFunc func(req *http.Request) (*http.Response, error)

// BaseResponse base response for all requests
type BaseResponse struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Client define API client
type Client struct {
	APIKey     string
	SecretKey  string
	BaseURL    string
	UserAgent  string
	HTTPClient *http.Client
	Debug      bool
	Logger     *log.Logger
	do         doFunc
}

func (c *Client) debug(format string, v ...interface{}) {
	if c.Debug {
		c.Logger.Printf(format, v...)
	}
}

func (c *Client) parseRequest(r *request, opts ...RequestOption) (err error) {
	// set request options from user
	for _, opt := range opts {
		opt(r)
	}
	err = r.validate()
	if err != nil {
		return err
	}

	fullURL := fmt.Sprintf("%s%s", c.BaseURL, r.endpoint)

	queryString := r.query.Encode()
	body := &bytes.Buffer{}
	bodyString := toJSON(r.form)
	header := http.Header{}

	if bodyString != "" {
		body = bytes.NewBufferString(bodyString)
		header.Set("content-type", "application/json")
	}

	if r.secType == secTypeAPIKey || r.secType == secTypeSigned {
		header.Set(accessTokenKey, c.APIKey)
	}

	expiry := fmt.Sprintf("%v", currentTimestamp()+60)
	if r.secType == secTypeSigned {
		header.Set(expiryKey, expiry)
	}

	if r.secType == secTypeSigned {
		raw := fmt.Sprintf("%s%s%s%s", r.endpoint, queryString, expiry, bodyString)
		signedString, err := c.signString(raw)
		if err != nil {
			return err
		}
		header.Set(signatureKey, signedString)
	}
	if queryString != "" {
		fullURL = fmt.Sprintf("%s?%s", fullURL, queryString)
	}
	c.debug("headers: %v", header)
	c.debug("full url: %s, body: %s", fullURL, bodyString)

	r.fullURL = fullURL
	r.header = header
	r.body = body
	return nil
}

func (c *Client) signString(raw string) (string, error) {
	c.debug("signed string: %v", raw)
	mac := hmac.New(sha256.New, []byte(c.SecretKey))
	_, err := mac.Write([]byte(raw))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(mac.Sum(nil)), nil
}

func toJSON(m interface{}) string {
	js, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	res := string(js)
	if res == "{}" {
		res = ""
	}
	return res
}

func (c *Client) callAPI(ctx context.Context, r *request, opts ...RequestOption) (data []byte, err error) {
	err = c.parseRequest(r, opts...)
	if err != nil {
		return []byte{}, err
	}
	req, err := http.NewRequest(r.method, r.fullURL, r.body)
	if err != nil {
		return []byte{}, err
	}
	req = req.WithContext(ctx)
	req.Header = r.header
	c.debug("request: %#v", req)
	f := c.do
	if f == nil {
		f = c.HTTPClient.Do
	}
	res, err := f(req)
	if err != nil {
		return []byte{}, err
	}
	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	defer func() {
		cerr := res.Body.Close()
		// Only overwrite the retured error if the original error was nil and an
		// error occurred while closing the body.
		if err == nil && cerr != nil {
			err = cerr
		}
	}()
	c.debug("response: %#v", res)
	c.debug("response body: %s", string(data))
	c.debug("response status code: %d", res.StatusCode)

	if res.StatusCode >= 400 {
		apiTempErr := new(apiTempError)
		e := json.Unmarshal(data, apiTempErr)
		if e != nil {
			c.debug("failed to unmarshal json: %s", e)
			return nil, e
		}
		apiErr := new(common.APIError)
		apiErr.Message = apiTempErr.Message
		apiErr.Code, e = strconv.ParseInt(apiTempErr.Code, 10, 64)
		if e != nil {
			c.debug("failed to parse int: %s", e)
			return nil, e
		}
		return nil, apiErr
	}
	return data, nil
}

type apiTempError struct {
	Code    string `json:"code"`
	Message string `json:"msg"`
}

// NewCreateOrderService init create order service
func (c *Client) NewCreateOrderService() *CreateOrderService {
	return &CreateOrderService{c: c}
}

// NewCreateReplaceOrderService init replace order service
func (c *Client) NewCreateReplaceOrderService() *CreateReplaceOrderService {
	return &CreateReplaceOrderService{c: c}
}

// NewCancelOrderService init cancel order service
func (c *Client) NewCancelOrderService() *CancelOrderService {
	return &CancelOrderService{c: c}
}

// NewQueryOrderService init query order service
func (c *Client) NewQueryOrderService() *QueryOrderService {
	return &QueryOrderService{c: c}
}

// NewPositionsLeverageService init positions leverage service
func (c *Client) NewPositionsLeverageService() *PositionsLeverageService {
	return &PositionsLeverageService{c: c}
}

// NewPositionsAssignService init positions assign service
func (c *Client) NewPositionsAssignService() *PositionsAssignService {
	return &PositionsAssignService{c: c}
}

// NewListOpenOrdersService init list open orders service
func (c *Client) NewListOpenOrdersService() *ListOpenOrdersService {
	return &ListOpenOrdersService{c: c}
}

// NewGetAccountPositionService init getting account position service
func (c *Client) NewGetAccountPositionService() *GetAccountPositionService {
	return &GetAccountPositionService{c: c}
}

// NewGetUserWalletService getting wallet service
func (c *Client) NewGetUserWalletService() *GetUserWalletService {
	return &GetUserWalletService{c: c}
}

// NewExchangeMarginService transfer funds between main wallet and margin account
func (c *Client) NewExchangeMarginService() *ExchangeMarginService {
	return &ExchangeMarginService{c: c}
}

// NewWsAuthService init starting auth ws service
func (c *Client) NewWsAuthService() *WsAuthService {
	return &WsAuthService{c: c}
}

// NewStartWsOrderBookService OrderBook service
func (c *Client) NewStartWsOrderBookService() *StartWsOrderBookService {
	return &StartWsOrderBookService{c: c}
}

// NewStartWsAOPService AOP service
func (c *Client) NewStartWsAOPService() *StartWsAOPService {
	return &StartWsAOPService{c: c}
}

// NewExchangeProductsService init exchange info service
func (c *Client) NewExchangeProductsService() *ExchangeProductsService {
	return &ExchangeProductsService{c: c}
}

// NewListPriceChangeStatsService init list prices change stats service
func (c *Client) NewListPriceChangeStatsService() *ListPriceChangeStatsService {
	return &ListPriceChangeStatsService{c: c}
}
