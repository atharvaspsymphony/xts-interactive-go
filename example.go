package main

import (
	"fmt"

	interactive "test/api"
)

var (
	clientID = ""
	Token    = ""
)

const (
	secretKey = "Epwe008@E6"
	appKey    = "df4f7ad5400d7b88fb9685"
	source    = "WebAPI"
	url       = "http://103.181.209.198:11091"
)

func main() {
	//login
	var loginPayload = interactive.LoginRequest{
		SecretKey: secretKey,
		AppKey:    appKey,
		Source:    "WebAPI",
	}
	login(loginPayload)
	interactive.Socket(url, Token, clientID)

	//Balance
	// getBalance()

	//Profile
	// getProfile()

	//ExchangeMessage
	// getExchangeMessage()

	//ExchangeStatus
	// getExchangeStaus()

	//Placeorder
	// var PlaceOrderPayload = interactive.PlaceOrderRequest{
	// 	ExchangeSegment:       "NSECM",
	// 	ExchangeInstrumentID:  "2029",
	// 	ProductType:           "NRML",
	// 	OrderType:             "MARKET",
	// 	OrderSide:             "BUY",
	// 	TimeInForce:           "DAY",
	// 	DisclosedQuantity:     0,
	// 	OrderQuantity:         15,
	// 	LimitPrice:            185,
	// 	StopPrice:             190,
	// 	OrderUniqueIdentifier: "123abc",
	// }
	// placeorder(PlaceOrderPayload)

	//ModifyOrder
	// var ModifyOrderPayload = interactive.ModifyOrderRequest{
	// 	AppOrderID:                "1210903429",
	// 	ModifiedProductType:       "NRML",
	// 	ModifiedOrderType:         "MARKET",
	// 	ModifiedOrderQuantity:     2,
	// 	ModifiedDisclosedQuantity: 0,
	// 	ModifiedLimitPrice:        2640,
	// 	ModifiedStopPrice:         2650,
	// 	ModifiedTimeInForce:       "DAY",
	// }
	// modifyorder(ModifyOrderPayload)

	//cancelOrder
	// var appOrderID = "1210903563"
	// cancelOrder(appOrderID)

	//OrderBook
	// getOrderBook()

	//OrderHistory
	// getOrdreHistory(appOrderID)

	//TradeBook
	// getTradeBook()

	//Holdings
	// getHoldings()

	//Positions
	// var dayOrNet = "DayWise"
	// getPositions(dayOrNet)

	//Squareoff
	// var SquareOffPayload = interactive.SquareOffRequest{
	// 	ExchangeSegment:               "NSECM",
	// 	ExchangeInstrumentID:          2029,
	// 	ProductType:                   "NRML",
	// 	SquareOffMode:                 "DayWise",
	// 	PositionSquareOffQuantityType: "ExactQty",
	// 	SquareOffQtyValue:             1,
	// 	BlockOrderSending:             "True",
	// 	CancelOrders:                  "True",
	// }
	// squareoff(SquareOffPayload)

	//CancelAll
	// var cancelallPayload = interactive.CancelAllRequest{
	// 	ExchangeSegment:      "NSECM",
	// 	ExchangeInstrumentID: "2885",
	// }
	// cancelallorders(cancelallPayload)

	//PositionConvert
	// var ConvertPositionPayload = interactive.ConvertPositionRequest{
	// 	ExchangeSegment:      "NSECM",
	// 	ExchangeInstrumentID: 11543,
	// 	OldProductType:       "NRML",
	// 	NewProductType:       "MIS",
	// 	IsDayWise:            true,
	// 	TargetQty:            1,
	// 	StatisticsLevel:      "ParentLevel",
	// 	IsInterOpPosition:    true,
	// }
	// convertPosition(ConvertPositionPayload)

	//Squareoffall
	// var squareoffallPayload = interactive.SquareOffAllRequest{
	// 	SquareOffMode: "NetWise",
	// 	ClientID:      clientID,
	// }
	// squareoffall(squareoffallPayload)

	//PlaceBracket Order
	// var bracketOrderPayload = interactive.BracketOrderRequest{
	// 	ClientID:              clientID,
	// 	OrderSide:             "BUY",
	// 	DisclosedQuantity:     0,
	// 	ExchangeSegment:       "NSECM",
	// 	ExchangeInstrumentID:  15083,
	// 	LimitPrice:            1495,
	// 	OrderType:             "LIMIT",
	// 	OrderQuantity:         25,
	// 	SquareOff:             10,
	// 	StopLossPrice:         10,
	// 	TrailingStopLoss:      0,
	// 	IsProOrder:            false,
	// 	OrderUniqueIdentifier: "adaniports_bracket",
	// }
	// placeBracketOrder(bracketOrderPayload)

	//ModifyBracket Order
	// var modifyBracketPayload = interactive.ModifyBracketOrderRequest{
	// 	ClientID:          clientID,
	// 	LimitPrice:        6790,
	// 	StopLossPrice:     10,
	// 	OrderQuantity:     2,
	// 	AppOrderID:        "1510901188",
	// 	ModifiedOrderType: "MARKET",
	// }
	// modifyBracketOrder(modifyBracketPayload)

	//CancelBracketOrder
	// var BoEntryOrderId = "1510900952"
	// cancelBO(BoEntryOrderId)

	//PlaceCoverOrder
	// var coverorderPayload = interactive.CoverOrderRequest{
	// 	ExchangeSegment:       "NSECM",
	// 	ExchangeInstrumentID:  3787,
	// 	OrderSide:             "BUY",
	// 	OrderQuantity:         15,
	// 	DisclosedQuantity:     0,
	// 	LimitPrice:            520,
	// 	StopPrice:             470,
	// 	OrderType:             "LIMIT",
	// 	OrderUniqueIdentifier: "wipro_co",
	// 	ClientID:              clientID,
	// }
	// placeCoverOrder(coverorderPayload)

	//ExitCoverOrder
	// var exitcoverpayload = interactive.ExitCoverOrderRequest{
	// 	AppOrderID: "1410903377",
	// }
	// exitcoverOrder(exitcoverpayload)

	//Logout
	// interactive.Logout()
}

func login(loginPayload interactive.LoginRequest) {
	response, err := interactive.Login(url, loginPayload)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	clientID = response.Result.UserID
	Token = response.Result.Token
	fmt.Println("Login Response-->", Token, clientID)
}

func getBalance() {
	response, err := interactive.Balance()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Balance Response-->", response)
}

func getProfile() {
	response, err := interactive.Profile(clientID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Profile Response-->", response)
}

func getExchangeMessage() {
	var exchangeSegment = "NSECM"
	response, err := interactive.ExchangeMessage(exchangeSegment)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Profile Response-->", response)
}

func getExchangeStaus() {
	response, err := interactive.ExchangeStatus(clientID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("ExchangeStatus Response-->", response)
}

func placeorder(PlaceOrderPayload interactive.PlaceOrderRequest) {
	response, err := interactive.PlaceOrder(PlaceOrderPayload)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("placeorder Response-->", response.Result.AppOrderID)
}

func modifyorder(ModifyOrderPayload interactive.ModifyOrderRequest) {
	response, err := interactive.ModifyOrder(ModifyOrderPayload)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("ModifyOrder Response-->", response.Result)
}

func cancelOrder(appOrderID string) {
	response, err := interactive.CancelOrder(appOrderID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("CancelOrder Response-->", response)
}

func getOrderBook() {
	response, err := interactive.OrderBook()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("OrderBook Response-->", response)
}

func getOrdreHistory(appOrderID string) {
	response, err := interactive.OrderHistory(appOrderID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("OrderHistory Response-->", response)
}

func getTradeBook() {
	response, err := interactive.TradeBook()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("TradeBook Response-->", response)
}

func getHoldings() {
	response, err := interactive.Holdings()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Holdings Response-->", response)
}

func getPositions(dayOrNet string) {
	response, err := interactive.Positions(dayOrNet)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Positions Response-->", response)
}

func squareoff(SquareOffPayload interactive.SquareOffRequest) {
	response, err := interactive.SquareOff(SquareOffPayload)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("squareoff Response-->", response.Result)
}

func cancelallorders(cancelallPayload interactive.CancelAllRequest) {
	response, err := interactive.CancelAllOrders(cancelallPayload)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("cancelall Response-->", response)
}

func convertPosition(ConvertPositionPayload interactive.ConvertPositionRequest) {
	response, err := interactive.ConvertPosition(ConvertPositionPayload)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("convertPosition Response-->", response)
}

func squareoffall(squareoffallPayload interactive.SquareOffAllRequest) {
	response, err := interactive.SquareOffAll(squareoffallPayload)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("squareoffall Response-->", response.Result)
}

func placeBracketOrder(bracketOrderPayload interactive.BracketOrderRequest) {
	response, err := interactive.PlaceBracketOrder(bracketOrderPayload)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Bracketorder Response-->", response.Result)
}

func modifyBracketOrder(modifyBracketPayload interactive.ModifyBracketOrderRequest) {
	response, err := interactive.ModifyBracketOrder(modifyBracketPayload)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("ModifyBracketorder Response-->", response)
}

func cancelBO(BoEntryOrderId string) {
	response, err := interactive.CancelBracketOrder(BoEntryOrderId)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("CancelBO Response-->", response)
}

func placeCoverOrder(coverorderPayload interactive.CoverOrderRequest) {
	response, err := interactive.PlaceCoverOrder(coverorderPayload)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("PlaceCoverOrder Response-->", response.Result)
}

func exitcoverOrder(exitcoverpayload interactive.ExitCoverOrderRequest) {
	response, err := interactive.ExitCoverOrder(exitcoverpayload)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("ExitCoverOrder Response-->", response)
}
