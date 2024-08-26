package interactive

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GenericHeader stores common headers.
var header = GenericHeader{
	contentType:   "application/json",
	Authorization: "",
}

// baseURL stores the base URL for the API.
var baseURL string = ""

// setHeaders sets common headers for a request.
func setHeaders(req *http.Request) {
	req.Header.Set("Content-Type", header.contentType)
	if header.Authorization != "" {
		req.Header.Set("Authorization", header.Authorization)
	}
}

func doRequest(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Handle non-200 status codes
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body) // Read body even on error to provide meaningful error messages
		return nil, fmt.Errorf("request failed with status: %s, body: %s", resp.Status, string(body))
	}

	return io.ReadAll(resp.Body)
}

// parseJSONResponse parses the JSON response into the provided interface.
func parseJSONResponse(body []byte, v interface{}) error {
	return json.Unmarshal(body, v)
}

// Login API.
func Login(url string, payload LoginRequest) (*LoginResponse, error) {
	baseURL = url

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", baseURL+interactiveRoutes["user.login"].(string), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	setHeaders(req)

	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var response LoginResponse
	if err := parseJSONResponse(body, &response); err != nil {
		return nil, err
	}

	header.Authorization = response.Result.Token
	return &response, nil
}

// Balance API.
func Balance() (*GenericResponse, error) {
	req, err := http.NewRequest("GET", baseURL+interactiveRoutes["user.balance"].(string), nil)
	if err != nil {
		return nil, err
	}
	setHeaders(req)

	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var response GenericResponse
	if err := parseJSONResponse(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// Profile API.
func Profile(clientID string) (*GenericResponse, error) {

	fullURL := fmt.Sprintf("%s?clientID=%s", baseURL+interactiveRoutes["user.profile"].(string), clientID)
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}
	setHeaders(req)

	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var response GenericResponse
	if err := parseJSONResponse(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// ExchangeMessage API.
func ExchangeMessage(exchangeSegment string) (*GenericResponse, error) {

	fmt.Println("exchangeSegment", exchangeSegment)
	fullURL := fmt.Sprintf("%s?exchangeSegment=%s", baseURL+interactiveRoutes["exchange.message"].(string), exchangeSegment)
	fmt.Println("fullURL-->", fullURL)
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}
	setHeaders(req)

	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var response GenericResponse
	if err := parseJSONResponse(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// ExchangeStatus API.
func ExchangeStatus(clientID string) (*GenericResponse, error) {

	fullURL := fmt.Sprintf("%s?userID=%s", baseURL+interactiveRoutes["exchange.status"].(string), clientID)
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}
	setHeaders(req)

	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var response GenericResponse
	if err := parseJSONResponse(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// PlaceOrder API
func PlaceOrder(PlaceOrderPayload PlaceOrderRequest) (*PlaceOrderResponse, error) {
	jsonData, err := json.Marshal(PlaceOrderPayload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", baseURL+interactiveRoutes["order"].(string), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	setHeaders(req)

	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var result PlaceOrderResponse
	if err := parseJSONResponse(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// ModifyOrder API
func ModifyOrder(ModifyOrderPayload ModifyOrderRequest) (*ModifyOrderResponse, error) {
	jsonData, err := json.Marshal(ModifyOrderPayload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", baseURL+interactiveRoutes["order"].(string), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	setHeaders(req)

	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var result ModifyOrderResponse
	if err := parseJSONResponse(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CancelOrder API.
func CancelOrder(appOrderID string) (*ModifyOrderResponse, error) {
	fullURL := fmt.Sprintf("%s?appOrderID=%s", baseURL+interactiveRoutes["order"].(string), appOrderID)
	req, err := http.NewRequest("DELETE", fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	setHeaders(req)

	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var response ModifyOrderResponse
	if err := parseJSONResponse(body, &response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	// Check for specific error in response
	if response.Type == "error" && response.Code == "e-request-0004" {
		return nil, fmt.Errorf("order not found: %s", response.Description)
	}

	return &response, nil
}

// OrderBook API.
func OrderBook() (*OrderBookResponse, error) {
	req, err := http.NewRequest("GET", baseURL+interactiveRoutes["order"].(string), nil)
	if err != nil {
		return nil, err
	}
	setHeaders(req)

	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var response OrderBookResponse
	if err := parseJSONResponse(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// OrderHistory API.
func OrderHistory(appOrderID string) (*OrderBookResponse, error) {
	fullURL := fmt.Sprintf("%s?appOrderID=%s", baseURL+interactiveRoutes["order"].(string), appOrderID)
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	setHeaders(req)

	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var response OrderBookResponse
	if err := parseJSONResponse(body, &response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	// Check for specific error in response
	if response.Type == "error" && response.Code == "e-request-0004" {
		return nil, fmt.Errorf("order not found: %s", response.Description)
	}

	return &response, nil
}

// TradeBook API.
func TradeBook() (*TradeBookResponse, error) {
	req, err := http.NewRequest("GET", baseURL+interactiveRoutes["order.tradebook"].(string), nil)
	if err != nil {
		return nil, err
	}
	setHeaders(req)

	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var response TradeBookResponse
	if err := parseJSONResponse(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// Holdings API.
func Holdings() (*GenericResponse, error) {
	req, err := http.NewRequest("GET", baseURL+interactiveRoutes["portfolio.holdings"].(string), nil)
	if err != nil {
		return nil, err
	}
	setHeaders(req)

	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var response GenericResponse
	if err := parseJSONResponse(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// Positions API
func Positions(dayOrNet string) (*PositionResponse, error) {

	fullURL := fmt.Sprintf("%s?dayOrNet=%s", baseURL+interactiveRoutes["portfolio.positions"].(string), dayOrNet)
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}
	setHeaders(req)

	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var response PositionResponse
	if err := parseJSONResponse(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// SquareOff API
func SquareOff(SquareOffPayload SquareOffRequest) (*SquareOffResponse, error) {
	jsonData, err := json.Marshal(SquareOffPayload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", baseURL+interactiveRoutes["portfolio.squareoff"].(string), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	setHeaders(req)

	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var result SquareOffResponse
	if err := parseJSONResponse(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CancelAll API
func CancelAllOrders(cancelallPayload CancelAllRequest) (*GenericResponse, error) {
	jsonData, err := json.Marshal(cancelallPayload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", baseURL+interactiveRoutes["order.cancelall"].(string), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	setHeaders(req)

	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var result GenericResponse
	if err := parseJSONResponse(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// ConvertPosition API
func ConvertPosition(ConvertPositionPayload ConvertPositionRequest) (*GenericResponse, error) {
	jsonData, err := json.Marshal(ConvertPositionPayload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", baseURL+interactiveRoutes["portfolio.positionConvert"].(string), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	setHeaders(req)

	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var result GenericResponse
	if err := parseJSONResponse(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func SquareOffAll(squareoffallPayload SquareOffAllRequest) (*GenericResponse, error) {
	jsonData, err := json.Marshal(squareoffallPayload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", baseURL+interactiveRoutes["portfolio.squareoffall"].(string), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	setHeaders(req)

	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var result GenericResponse
	if err := parseJSONResponse(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// PlaceBracketOrder API
func PlaceBracketOrder(bracketOrderPayload BracketOrderRequest) (*PlaceOrderResponse, error) {
	jsonData, err := json.Marshal(bracketOrderPayload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", baseURL+interactiveRoutes["order.bracket"].(string), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	setHeaders(req)

	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var result PlaceOrderResponse
	if err := parseJSONResponse(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// ModifyBracketOrder API
func ModifyBracketOrder(modifyBracketPayload ModifyBracketOrderRequest) (*ModifyOrderResponse, error) {
	jsonData, err := json.Marshal(modifyBracketPayload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", baseURL+interactiveRoutes["order.bracket"].(string), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	setHeaders(req)

	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var result ModifyOrderResponse
	if err := parseJSONResponse(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CancelBracketOrder API.
func CancelBracketOrder(appOrderID string) (*ModifyOrderResponse, error) {
	fullURL := fmt.Sprintf("%s?boEntryOrderId=%s", baseURL+interactiveRoutes["order.bracket"].(string), appOrderID)
	req, err := http.NewRequest("DELETE", fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	setHeaders(req)

	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var response ModifyOrderResponse
	if err := parseJSONResponse(body, &response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	// Check for specific error in response
	if response.Type == "error" && response.Code == "e-request-0004" {
		return nil, fmt.Errorf("order not found: %s", response.Description)
	}

	return &response, nil
}

// PlaceCoverOrder API
func PlaceCoverOrder(coverorderPayload CoverOrderRequest) (*CoverOrderResponse, error) {
	jsonData, err := json.Marshal(coverorderPayload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", baseURL+interactiveRoutes["order.cover"].(string), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	setHeaders(req)

	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var result CoverOrderResponse
	if err := parseJSONResponse(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// ExitCoverOrder API
func ExitCoverOrder(exitcoverpayload ExitCoverOrderRequest) (*ModifyOrderResponse, error) {
	jsonData, err := json.Marshal(exitcoverpayload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", baseURL+interactiveRoutes["order.cover"].(string), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	setHeaders(req)

	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var result ModifyOrderResponse
	if err := parseJSONResponse(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// Logout API
func Logout() (*GenericResponse, error) {
	req, err := http.NewRequest("DELETE", baseURL+interactiveRoutes["auth.logout"].(string), nil)
	if err != nil {
		return nil, err
	}
	setHeaders(req)
	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var result GenericResponse
	if err := parseJSONResponse(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
