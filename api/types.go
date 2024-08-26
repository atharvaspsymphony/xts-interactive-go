package interactive

// GenericResponse represents a generic structure of the response payload
type GenericResponse struct {
	Type        string      `json:"type"`
	Code        string      `json:"code"`
	Description string      `json:"description"`
	Result      interface{} `json:"result"`
}

type GenericHeader struct {
	contentType   string
	Authorization string
}

// LoginRequest represents the structure of the login request payload
type LoginRequest struct {
	SecretKey string `json:"secretKey"`
	AppKey    string `json:"appKey"`
	Source    string `json:"source"`
}

type LoginResponse struct {
	GenericResponse
	Result LoginResult `json:"result"`
}

// Result represents the structure of the result object in the response
type LoginResult struct {
	Enums                interface{} `json:"enums"`
	ClientCodes          []string    `json:"clientCodes"`
	ExchangeSegmentArray []struct{}  `json:"exchangeSegmentArray"`
	Token                string      `json:"token"`
	UserID               string      `json:"userID"`
	IsInvestorClient     bool        `json:"isInvestorClient"`
	IsOneTouchUser       bool        `json:"isOneTouchUser"`
}

type PlaceOrderRequest struct {
	DisclosedQuantity     int     `json:"disclosedQuantity"`
	ExchangeSegment       string  `json:"exchangeSegment"`
	OrderType             string  `json:"orderType"`
	StopPrice             float64 `json:"stopPrice"`
	LimitPrice            float64 `json:"limitPrice"`
	OrderUniqueIdentifier string  `json:"orderUniqueIdentifier"`
	ExchangeInstrumentID  string  `json:"exchangeInstrumentID"`
	OrderSide             string  `json:"orderSide"`
	TimeInForce           string  `json:"timeInForce"`
	ProductType           string  `json:"productType"`
	OrderQuantity         int     `json:"orderQuantity"`
}

type PlaceOrderResult struct {
	AppOrderID            int    `json:"AppOrderID"`
	OrderUniqueIdentifier string `json:"OrderUniqueIdentifier"`
	ClientID              string `json:"ClientID"`
}

type PlaceOrderResponse struct {
	Type        string           `json:"type"`
	Code        string           `json:"code"`
	Description string           `json:"description"`
	Result      PlaceOrderResult `json:"result"`
}

type ModifyOrderRequest struct {
	AppOrderID                string  `json:"appOrderID"`
	ModifiedProductType       string  `json:"modifiedProductType"`
	ModifiedOrderType         string  `json:"modifiedOrderType"`
	ModifiedOrderQuantity     int     `json:"modifiedOrderQuantity"`
	ModifiedDisclosedQuantity int     `json:"modifiedDisclosedQuantity"`
	ModifiedLimitPrice        float64 `json:"modifiedLimitPrice"`
	ModifiedStopPrice         float64 `json:"modifiedStopPrice"`
	ModifiedTimeInForce       string  `json:"modifiedTimeInForce"`
	OrderUniqueIdentifier     string  `json:"orderUniqueIdentifier"`
}

type ModifyOrderResponse struct {
	Type        string `json:"type"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Result      struct {
		AppOrderID int    `json:"AppOrderID"`
		ClientID   string `json:"ClientID"`
	}
}

type OrderBookResult struct {
	LoginID                 string  `json:"LoginID"`
	ClientID                string  `json:"ClientID"`
	AppOrderID              int     `json:"AppOrderID"`
	OrderReferenceID        string  `json:"OrderReferenceID"`
	GeneratedBy             string  `json:"GeneratedBy"`
	ExchangeOrderID         string  `json:"ExchangeOrderID"`
	OrderCategoryType       string  `json:"OrderCategoryType"`
	ExchangeSegment         string  `json:"ExchangeSegment"`
	ExchangeInstrumentID    int     `json:"ExchangeInstrumentID"`
	OrderSide               string  `json:"OrderSide"`
	OrderType               string  `json:"OrderType"`
	ProductType             string  `json:"ProductType"`
	TimeInForce             string  `json:"TimeInForce"`
	OrderPrice              float64 `json:"OrderPrice"`
	OrderQuantity           int     `json:"OrderQuantity"`
	OrderStopPrice          float64 `json:"OrderStopPrice"`
	OrderStatus             string  `json:"OrderStatus"`
	OrderAverageTradedPrice string  `json:"OrderAverageTradedPrice"`
	LeavesQuantity          int     `json:"LeavesQuantity"`
	CumulativeQuantity      int     `json:"CumulativeQuantity"`
	OrderDisclosedQuantity  int     `json:"OrderDisclosedQuantity"`
	OrderGeneratedDateTime  string  `json:"OrderGeneratedDateTime"`
	ExchangeTransactTime    string  `json:"ExchangeTransactTime"`
	TradingSymbol           string  `json:"TradingSymbol"`
	LastUpdateDateTime      string  `json:"LastUpdateDateTime"`
	OrderExpiryDate         string  `json:"OrderExpiryDate"`
	CancelRejectReason      string  `json:"CancelRejectReason"`
	OrderUniqueIdentifier   string  `json:"OrderUniqueIdentifier"`
	OrderLegStatus          string  `json:"OrderLegStatus"`
	BoLegDetails            int     `json:"BoLegDetails"`
	IsSpread                bool    `json:"IsSpread"`
	BoEntryOrderId          string  `json:"BoEntryOrderId"`
	ApiOrderSource          string  `json:"ApiOrderSource"`
	MessageCode             int     `json:"MessageCode"`
	MessageVersion          int     `json:"MessageVersion"`
	TokenID                 int     `json:"TokenID"`
	ApplicationType         int     `json:"ApplicationType"`
	SequenceNumber          int     `json:"SequenceNumber"`
}

type OrderBookResponse struct {
	Type        string            `json:"type"`
	Code        string            `json:"code"`
	Description string            `json:"description"`
	Result      []OrderBookResult `json:"result"`
}

type TradeBookResult struct {
	LoginID                   string  `json:"LoginID"`
	ClientID                  string  `json:"ClientID"`
	AppOrderID                int64   `json:"AppOrderID"`
	OrderReferenceID          string  `json:"OrderReferenceID"`
	GeneratedBy               string  `json:"GeneratedBy"`
	ExchangeOrderID           string  `json:"ExchangeOrderID"`
	OrderCategoryType         string  `json:"OrderCategoryType"`
	ExchangeSegment           string  `json:"ExchangeSegment"`
	ExchangeInstrumentID      int64   `json:"ExchangeInstrumentID"`
	OrderSide                 string  `json:"OrderSide"`
	OrderType                 string  `json:"OrderType"`
	ProductType               string  `json:"ProductType"`
	TimeInForce               string  `json:"TimeInForce"`
	OrderPrice                float64 `json:"OrderPrice"`
	OrderQuantity             int     `json:"OrderQuantity"`
	OrderStopPrice            float64 `json:"OrderStopPrice"`
	OrderStatus               string  `json:"OrderStatus"`
	OrderAverageTradedPrice   string  `json:"OrderAverageTradedPrice"`
	LeavesQuantity            int     `json:"LeavesQuantity"`
	CumulativeQuantity        int     `json:"CumulativeQuantity"`
	OrderDisclosedQuantity    int     `json:"OrderDisclosedQuantity"`
	OrderGeneratedDateTime    string  `json:"OrderGeneratedDateTime"`
	ExchangeTransactTime      string  `json:"ExchangeTransactTime"`
	LastUpdateDateTime        string  `json:"LastUpdateDateTime"`
	OrderUniqueIdentifier     string  `json:"OrderUniqueIdentifier"`
	OrderLegStatus            string  `json:"OrderLegStatus"`
	LastTradedPrice           float64 `json:"LastTradedPrice"`
	LastTradedQuantity        int     `json:"LastTradedQuantity"`
	LastExecutionTransactTime string  `json:"LastExecutionTransactTime"`
	ExecutionID               string  `json:"ExecutionID"`
	TradingSymbol             string  `json:"TradingSymbol"`
	ExecutionReportIndex      int     `json:"ExecutionReportIndex"`
	ApiOrderSource            string  `json:"ApiOrderSource"`
	IsSpread                  bool    `json:"IsSpread"`
	MessageCode               int     `json:"MessageCode"`
	MessageVersion            int     `json:"MessageVersion"`
	TokenID                   int     `json:"TokenID"`
	ApplicationType           int     `json:"ApplicationType"`
	SequenceNumber            int     `json:"SequenceNumber"`
}

type TradeBookResponse struct {
	Type        string            `json:"type"`
	Code        string            `json:"code"`
	Description string            `json:"description"`
	Result      []TradeBookResult `json:"result"`
}

type PositionResponse struct {
	Type        string `json:"type"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Result      struct {
		PositionList []struct {
			AccountID                       string   `json:"AccountID"`
			TradingSymbol                   string   `json:"TradingSymbol"`
			ExchangeSegment                 string   `json:"ExchangeSegment"`
			ExchangeInstrumentID            string   `json:"ExchangeInstrumentId"`
			ProductType                     string   `json:"ProductType"`
			Marketlot                       string   `json:"Marketlot"`
			Multiplier                      string   `json:"Multiplier"`
			BuyAveragePrice                 string   `json:"BuyAveragePrice"`
			SellAveragePrice                string   `json:"SellAveragePrice"`
			OpenBuyQuantity                 string   `json:"OpenBuyQuantity"`
			OpenSellQuantity                string   `json:"OpenSellQuantity"`
			Quantity                        string   `json:"Quantity"`
			BuyAmount                       string   `json:"BuyAmount"`
			SellAmount                      string   `json:"SellAmount"`
			NetAmount                       string   `json:"NetAmount"`
			UnrealizedMTM                   string   `json:"UnrealizedMTM"`
			RealizedMTM                     string   `json:"RealizedMTM"`
			MTM                             string   `json:"MTM"`
			BEP                             string   `json:"BEP"`
			SumOfTradedQuantityAndPriceBuy  string   `json:"SumOfTradedQuantityAndPriceBuy"`
			SumOfTradedQuantityAndPriceSell string   `json:"SumOfTradedQuantityAndPriceSell"`
			StatisticsLevel                 string   `json:"StatisticsLevel"`
			IsInterOpPosition               bool     `json:"IsInterOpPosition"`
			ChildPositions                  struct{} `json:"childPositions"`
			MessageCode                     int      `json:"MessageCode"`
			MessageVersion                  int      `json:"MessageVersion"`
			TokenID                         int      `json:"TokenID"`
			ApplicationType                 int      `json:"ApplicationType"`
			SequenceNumber                  int      `json:"SequenceNumber"`
		} `json:"positionList"`
	} `json:"result"`
}

type SquareOffRequest struct {
	ExchangeSegment               string `json:"exchangeSegment"`
	ExchangeInstrumentID          int    `json:"exchangeInstrumentID"`
	ProductType                   string `json:"productType"`
	SquareOffMode                 string `json:"squareoffMode"`
	PositionSquareOffQuantityType string `json:"positionSquareOffQuantityType"`
	SquareOffQtyValue             int    `json:"squareOffQtyValue"`
	BlockOrderSending             string `json:"blockOrderSending"`
	CancelOrders                  string `json:"cancelOrders"`
}

type SquareOffResponse struct {
	Type        string `json:"type"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Result      struct {
		UserID                        string `json:"UserID"`
		SquareOffMode                 string `json:"SqureOffmode"`
		AccountID                     string `json:"AccountID"`
		ExchangeSegment               string `json:"ExchangeSegment"`
		InstrumentID                  string `json:"InstrumentID"`
		RMSProduct                    string `json:"RMSProduct"`
		PositionSquareOffQuantityType string `json:"PositionSquareOffQuantityType"`
		SquareOffQtyValue             string `json:"SquareOffQtyValue"`
		SquareOffPriceValue           string `json:"SquareOffPriceValue"`
		IsOperationSuccess            bool   `json:"IsOperationSuccess"`
	} `json:"result"`
}

type CancelAllRequest struct {
	ExchangeSegment      string `json:"exchangeSegment"`
	ExchangeInstrumentID string `json:"exchangeInstrumentID"`
}

type ConvertPositionRequest struct {
	ExchangeSegment      string `json:"exchangeSegment"`
	ExchangeInstrumentID int    `json:"exchangeInstrumentID"`
	OldProductType       string `json:"oldProductType"`
	NewProductType       string `json:"newProductType"`
	IsDayWise            bool   `json:"isDayWise"`
	TargetQty            int    `json:"targetQty"`
	StatisticsLevel      string `json:"statisticsLevel"`
	IsInterOpPosition    bool   `json:"isInterOpPosition"`
}

type SquareOffAllRequest struct {
	SquareOffMode string `json:"SquareOffMode"`
	ClientID      string `json:"ClientID"`
}

type BracketOrderRequest struct {
	ClientID              string  `json:"clientID"`
	OrderSide             string  `json:"orderSide"`
	DisclosedQuantity     int     `json:"disclosedQuantity"`
	ExchangeSegment       string  `json:"exchangeSegment"`
	ExchangeInstrumentID  int     `json:"exchangeInstrumentID"`
	LimitPrice            float64 `json:"limitPrice"`
	OrderType             string  `json:"orderType"`
	OrderQuantity         int     `json:"orderQuantity"`
	SquareOff             float64 `json:"squarOff"`
	StopLossPrice         float64 `json:"stopLossPrice"`
	TrailingStopLoss      float64 `json:"trailingStoploss"`
	IsProOrder            bool    `json:"isProOrder"`
	OrderUniqueIdentifier string  `json:"orderUniqueIdentifier"`
}

type ModifyBracketOrderRequest struct {
	ClientID          string  `json:"clientID"`
	LimitPrice        float64 `json:"limitPrice"`
	StopLossPrice     float64 `json:"stopLossPrice"`
	OrderQuantity     int     `json:"orderQuantity"`
	AppOrderID        string  `json:"appOrderID"`
	ModifiedOrderType string  `json:"modifiedOrderType"`
}

type CoverOrderRequest struct {
	ExchangeSegment       string  `json:"exchangeSegment"`
	ExchangeInstrumentID  int     `json:"exchangeInstrumentID"`
	OrderSide             string  `json:"orderSide"`
	OrderQuantity         int     `json:"orderQuantity"`
	DisclosedQuantity     int     `json:"disclosedQuantity"`
	LimitPrice            float64 `json:"limitPrice"`
	StopPrice             float64 `json:"stopPrice"`
	OrderType             string  `json:"orderType"`
	OrderUniqueIdentifier string  `json:"orderUniqueIdentifier"`
	ClientID              string  `json:"clientID"`
}

type CoverOrderResponse struct {
	Type        string `json:"type"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Result      struct {
		EntryAppOrderID       int    `json:"EntryAppOrderID"`
		ExitAppOrderID        int    `json:"ExitAppOrderID"`
		OrderUniqueIdentifier string `json:"OrderUniqueIdentifier"`
		ClientID              string `json:"ClientID"`
	} `json:"result"`
}

type ExitCoverOrderRequest struct {
	AppOrderID string `json:"appOrderID"`
}
