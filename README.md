### go-phemex

A Golang SDK for [phemex](https://www.phemex.com) API.

[![Build Status](https://travis-ci.org/Krisa/go-phemex.svg?branch=master)](https://travis-ci.org/Krisa/go-phemex)
[![GoDoc](https://godoc.org/github.com/Krisa/go-phemex?status.svg)](https://godoc.org/github.com/Krisa/go-phemex)
[![Go Report Card](https://goreportcard.com/badge/github.com/Krisa/go-phemex)](https://goreportcard.com/report/github.com/Krisa/go-phemex)
[![codecov](https://codecov.io/gh/Krisa/go-phemex/branch/master/graph/badge.svg)](https://codecov.io/gh/Krisa/go-phemex)

Most REST APIs listed in [phemex API document](https://github.com/phemex/phemex-api-docs/blob/master/Public-API-en.md) are implemented, as well as the AOP websocket APIs.

For best compatibility, please use Go >= 1.8.

Make sure you have read phemex API document before continuing.

### API List

Name | Description | Status
------------ | ------------ | ------------
[rest-api.md](https://github.com/phemex/phemex-api-docs/blob/master/Public-API-en.md) | Details on the Rest API 

### Installation

```shell
go get github.com/Krisa/go-phemex
```

### Importing

```golang
import (
    "github.com/Krisa/go-phemex"
)
```

### Documentation

[![GoDoc](https://godoc.org/github.com/Krisa/go-phemex?status.svg)](https://godoc.org/github.com/Krisa/go-phemex)

### REST API

#### Setup

Init client for API services. Get APIKey/SecretKey from your phemex account.

```golang
var (
    apiKey = "your api key"
    secretKey = "your secret key"
)
client := phemex.NewClient(apiKey, secretKey)
```

A service instance stands for a REST API endpoint and is initialized by client.NewXXXService function.

Simply call API in chain style. Call Do() in the end to send HTTP request.

#### Create Order

```golang
order, err := client.NewCreateOrderService().Symbol("BTCUSD").
        Side(phemex.SideTypeBuy).Type(phemex.OrderTypeLimit).
        TimeInForce(phemex.TimeInForceTypeGTC).Quantity("5").
        Price("1000").Do(context.Background())
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(order)

```

#### Cancel Order

```golang
_, err := client.NewCancelOrderService().Symbol("BTCUSD").
    OrderID(4432844).Do(context.Background())
if err != nil {
    fmt.Println(err)
    return
}
```

#### List Open Orders

```golang
openOrders, err := client.NewListOpenOrdersService().Symbol("BTCUSD").
    Do(context.Background())
if err != nil {
    fmt.Println(err)
    return
}
for _, o := range openOrders {
    fmt.Println(o)
}
```

#### Get Account

```golang
res, err := client.NewGetAccountPositionService().Currency("BTC").Do(context.Background())
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(res)
```

### Websocket

You initiate PhemexClient the same way

#### User Data

```golang
wsHandler := func(message interface{}) {
    switch message := message.(type) {
    case *phemex.WsAOP:
        // snapshots / increments
    case *phemex.WsPositionInfo:
        // when a position is active
    case *phemex.WsError:
        // on connection
    }
}

errHandler := func(err error) {
    // initiate reconnection with `once.Do...`
}

auth := PhemexClient.NewWsAuthService()

if test {
    auth = auth.URL("wss://testnet.phemex.com/ws")
}

c, err := auth.Do(context.Background())
// err handling

err = PhemexClient.NewStartWsAOPService().SetID(1).Do(c, wsHandler, errHandler)
// err handling
```
