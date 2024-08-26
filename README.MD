# xts-interactive-GO

Welcome to the XTS-Interactive-API SDK version 3, This repository is Official Node.js library for Symphony XTS-Interactive-API.

API Documentation for XTS-trading API can be found in the below link.

https://symphonyfintech.com/xts-trading-front-end-api/

The XTS Trading API provides developer, data-scientist, financial analyst and investor the functionality necessary to create automated trading strategies, as well as other trading related applications with support of XTS OEMS hosted by Financial Institutions to trade with Indian electronic exchanges.

With the use of the socket.io library, the API has streaming capability and will push data notifications in a JSON format. Your application can place orders and receive real-time trade notification.

There is also an examples folder available which illustrates how to create a connection to XTS OEMS hosted by Brokers to subscribe to real-time events. Please request for apikeys with Symphony Fintech developer support team to start integrating your application with XTS OEMS.

## Installation
Clone the Github repo in your working dir
```bash
https://github.com/atharvaspsymphony/xts-interactive-go/tree/master
```
## Usage
Access the /api dir directly.

```go
import (
	"fmt"
	interactive "test/api"
)
```
Initailize all the required constants & variables. 

## Detailed explanation of API

### Login
call the login API to generate the token
POST /auth/login
```go
var loginPayload = interactive.LoginRequest{
	SecretKey: secretKey,
	AppKey:    appKey,
	Source:    "WebAPI",
}
response, err := interactive.Login(url, loginPayload)
```
Once the token is generated you can call any api provided in the documentation.

### Balance
GET /user/balance
```go
response, err := interactive.Balance()
```

### Profile
GET /user/profile
```go
response, err := interactive.Profile(clientID)
```

### ExchangeStatus
GET /status/exchange
```go
response, err := interactive.ExchangeStatus(clientID)
```

### ExchangeMessage
GET /message/exchange
```go
var exchangeSegment = "NSECM"
response, err := interactive.ExchangeMessage(exchangeSegment)
```

### Holding
GET /portfolio/holdings
```go
response, err := interactive.Holdings()
```

### Positions
GET /portfolio/positions
```go
response, err := interactive.Positions(dayOrNet)
```

### PositionConvert
PUT /portfolio/positions/convert
```go
var ConvertPositionPayload = interactive.ConvertPositionRequest{
	ExchangeSegment:      "NSECM",
	ExchangeInstrumentID: 11543,
	OldProductType:       "NRML",
	NewProductType:       "MIS",
	IsDayWise:            true,
	TargetQty:            1,
	StatisticsLevel:      "ParentLevel",
	IsInterOpPosition:    true,
}
response, err := interactive.ConvertPosition(ConvertPositionPayload)
```

### SquareOff
PUT /portfolio/positions/squareoff
```go
var squareoffallPayload = interactive.SquareOffAllRequest{
    SquareOffMode: "NetWise",
    ClientID:      clientID,
}
response, err := interactive.SquareOff(SquareOffPayload)
```

### SquareOffAll
PUT /portfolio/positions/squareoffall
```go
var squareoffallPayload = interactive.SquareOffAllRequest{
    SquareOffMode: "NetWise",
    ClientID:      clientID,
}
response, err := interactive.SquareOffAll(squareoffallPayload)
```

### Place Order
POST /orders
```go
var PlaceOrderPayload = interactive.PlaceOrderRequest{
    ExchangeSegment:       "NSECM",
    ExchangeInstrumentID:  "2029",
    ProductType:           "NRML",
    OrderType:             "MARKET",
    OrderSide:             "BUY",
    TimeInForce:           "DAY",
    DisclosedQuantity:     0,
    OrderQuantity:         15,
    LimitPrice:            185,
    StopPrice:             190,
    OrderUniqueIdentifier: "123abc",
}
response, err := interactive.PlaceOrder(PlaceOrderPayload)
```

### Modify Order
PUT /orders
```go
var ModifyOrderPayload = interactive.ModifyOrderRequest{
    AppOrderID:                "1210903429",
    ModifiedProductType:       "NRML",
    ModifiedOrderType:         "MARKET",
    ModifiedOrderQuantity:     2,
    ModifiedDisclosedQuantity: 0,
    ModifiedLimitPrice:        2640,
    ModifiedStopPrice:         2650,
    ModifiedTimeInForce:       "DAY",
}
response, err := interactive.ModifyOrder(ModifyOrderPayload)
```

### Cancel Order
DELETE /orders
```go
var appOrderID = "1210903563"
response, err := interactive.CancelOrder(appOrderID)
```

### CancelAll Order
POST /orders/cancelall
```go
var cancelallPayload = interactive.CancelAllRequest{
    ExchangeSegment:      "NSECM",
    ExchangeInstrumentID: "2885",
}
response, err := interactive.CancelAllOrders(cancelallPayload)
```

### Place Bracket Order
POST /orders/bracket
```go
var bracketOrderPayload = interactive.BracketOrderRequest{
    ClientID:              clientID,
    OrderSide:             "BUY",
    DisclosedQuantity:     0,
    ExchangeSegment:       "NSECM",
    ExchangeInstrumentID:  15083,
    LimitPrice:            1495,
    OrderType:             "LIMIT",
    OrderQuantity:         25,
    SquareOff:             10,
    StopLossPrice:         10,
    TrailingStopLoss:      0,
    IsProOrder:            false,
    OrderUniqueIdentifier: "adaniports_bracket",
}
response, err := interactive.PlaceBracketOrder(bracketOrderPayload)
```

### Modify Bracket Order
PUT /orders/bracket
```go
var modifyBracketPayload = interactive.ModifyBracketOrderRequest{
    ClientID:          clientID,
    LimitPrice:        6790,
    StopLossPrice:     10,
    OrderQuantity:     2,
    AppOrderID:        "1510901188",
    ModifiedOrderType: "MARKET",
}
response, err := interactive.ModifyBracketOrder(modifyBracketPayload)
```

### Cancel Bracket Order
DELETE /orders/bracket
```go
var BoEntryOrderId = "1510900952"
response, err := interactive.CancelBracketOrder(BoEntryOrderId)
```

### Place Cover Order
POST /orders/cover
```go
var coverorderPayload = interactive.CoverOrderRequest{
    ExchangeSegment:       "NSECM",
    ExchangeInstrumentID:  3787,
    OrderSide:             "BUY",
    OrderQuantity:         15,
    DisclosedQuantity:     0,
    LimitPrice:            520,
    StopPrice:             470,
    OrderType:             "LIMIT",
    OrderUniqueIdentifier: "wipro_co",
    ClientID:              clientID,
}
response, err := interactive.PlaceCoverOrder(coverorderPayload)
```

### Exit Cover Order
PUT /orders/cover
```go
var exitcoverpayload = interactive.ExitCoverOrderRequest{
    AppOrderID: "1410903377",
}
response, err := interactive.ExitCoverOrder(exitcoverpayload)
```

### Order Book
GET /orders
```go
response, err := interactive.OrderBook()
```

### Order History
GET /orders
```go
response, err := interactive.OrderHistory(appOrderID)
```

### Trade Book
GET /orders/trades
```go
response, err := interactive.TradeBook()
```

### Logout
DELETE /user/session
```go
response,err = interactive.Logout()
```

## Instantiating the XtsInteractive Websocket
This component provides functionality to access the socket related events. All real-time events can be registered. After token is generated, you can access the socket component and instantiate the socket.

Note:- XTS Interactive WebSocket is based on "https://socket.io/" library. This library is available in most of the programming languages. In this package code third-party socket-io library is used to make client connection for socket. This is just an example to make connection for websocket.

```go
response, err := interactive.Login(url, loginPayload)
interactive.Socket(url, Token, clientID)
```