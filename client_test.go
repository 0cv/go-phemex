package phemex

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type baseTestSuite struct {
	suite.Suite
	client    *mockedClient
	apiKey    string
	secretKey string
}

func (s *baseTestSuite) r() *require.Assertions {
	return s.Require()
}

func (s *baseTestSuite) SetupTest() {
	s.apiKey = "dummyAPIKey"
	s.secretKey = "dummySecretKey"
	s.client = newMockedClient(s.apiKey, s.secretKey)
}

func (s *baseTestSuite) mockDo(data []byte, err error, statusCode ...int) {
	s.client.Client.do = s.client.do
	code := http.StatusOK
	if len(statusCode) > 0 {
		code = statusCode[0]
	}
	s.client.On("do", anyHTTPRequest()).Return(newHTTPResponse(data, code), err)
}

func (s *baseTestSuite) assertDo() {
	s.client.AssertCalled(s.T(), "do", anyHTTPRequest())
}

func (s *baseTestSuite) assertReq(f func(r *request)) {
	s.client.assertReq = f
}

func (s *baseTestSuite) assertRequestEqual(e, a *request) {
	s.assertHeadersEquals(e.header, a.header)
	s.Equal(e.form, a.form)
}

func (s *baseTestSuite) assertHeadersEquals(e, a http.Header) {
	var eKeys, aKeys []string
	for k := range e {
		eKeys = append(eKeys, k)
	}
	for k := range a {
		aKeys = append(aKeys, k)
	}
	r := s.r()
	r.Len(aKeys, len(eKeys))
	for k := range a {
		switch k {
		case accessTokenKey, expiryKey, signatureKey:
			r.NotEmpty(a.Get(k))
			continue
		}
		r.Equal(e.Get(k), a.Get(k), k)
	}
}

func anythingOfType(t string) mock.AnythingOfTypeArgument {
	return mock.AnythingOfType(t)
}

func newContext() context.Context {
	return context.Background()
}

func anyHTTPRequest() mock.AnythingOfTypeArgument {
	return anythingOfType("*http.Request")
}

func newHTTPResponse(data []byte, statusCode int) *http.Response {
	return &http.Response{
		Body:       ioutil.NopCloser(bytes.NewBuffer(data)),
		StatusCode: statusCode,
	}
}

func newRequest() *request {
	r := &request{
		query: url.Values{},
		form:  map[string]interface{}{},
	}
	return r
}

func newSignedRequest() *request {
	return newRequest().setParams(params{
		accessTokenKey: "",
		expiryKey:      "",
		signatureKey:   "",
	})
}

type assertReqFunc func(r *request)

type mockedClient struct {
	mock.Mock
	*Client
	assertReq assertReqFunc
}

func newMockedClient(apiKey, secretKey string) *mockedClient {
	m := new(mockedClient)
	m.Client = NewClient(apiKey, secretKey)
	return m
}

func (m *mockedClient) do(req *http.Request) (*http.Response, error) {
	if m.assertReq != nil {
		r := newRequest()
		r.query = req.URL.Query()
		if req.Body != nil && req.ContentLength > 0 {
			bs := make([]byte, req.ContentLength)
			for {
				n, _ := req.Body.Read(bs)
				if n == 0 {
					break
				}
			}
			// panic(fmt.Sprintf("CONTENT => %[1]v --", req))
			var form map[string]interface{}
			err := json.Unmarshal(bs, &form)
			if err != nil {
				panic(err)
			}
			r.form = form
		}
		m.assertReq(r)
	}
	args := m.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}
