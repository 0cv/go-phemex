package phemex

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type basePositionTestSuite struct {
	baseTestSuite
}

type positionServiceTestSuite struct {
	basePositionTestSuite
}

func TestPositionService(t *testing.T) {
	suite.Run(t, new(positionServiceTestSuite))
}

func (s *positionServiceTestSuite) TestPositionLeverageService() {
	data := []byte(`{
		"code": 0,
		"msg": "OK"
	}`)

	s.mockDo(data, nil)
	defer s.assertDo()
	symbol := "BTCUSD"
	leverage := int64(10)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":   symbol,
			"leverage": leverage,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewPositionsLeverageService().Symbol(symbol).Leverage(leverage).Do(newContext())
	s.r().NoError(err)
	e := &BaseResponse{
		Code: 0,
		Msg:  "OK",
	}
	s.assertLeverageEquals(e, res)
}

func (s *positionServiceTestSuite) assertLeverageEquals(e, a *BaseResponse) {
	r := s.r()
	r.Equal(e.Code, a.Code, "Code")
	r.Equal(e.Msg, a.Msg, "Msg")
}

func (s *positionServiceTestSuite) TestPositionAssignService() {
	data := []byte(`{
		"code": 0,
		"msg": "OK"
	}`)

	s.mockDo(data, nil)
	defer s.assertDo()
	symbol := "BTCUSD"
	posBalance := float64(10)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":     symbol,
			"posBalance": posBalance,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewPositionsAssignService().Symbol(symbol).PosBalance(posBalance).Do(newContext())
	s.r().NoError(err)
	e := &BaseResponse{
		Code: 0,
		Msg:  "OK",
	}
	s.assertPositionAssignEquals(e, res)
}

func (s *positionServiceTestSuite) assertPositionAssignEquals(e, a *BaseResponse) {
	r := s.r()
	r.Equal(e.Code, a.Code, "Code")
	r.Equal(e.Msg, a.Msg, "Msg")
}
