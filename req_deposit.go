package go_neteller

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// 下单(充值/提现是同一个接口)
func (cli *Client) Deposit(req NetellerPaymentReq) (*NetellerPaymentRsp, error) {
	req.OutType = int(Deposit)
	return cli.CreatePaymentHandle(&req)
}

// 预充值/预提现:都是走这里
func (cli *Client) CreatePaymentHandle(req *NetellerPaymentReq) (*NetellerPaymentRsp, error) {
	// Prepare request data
	amountInCents := req.Amount * 100 //.Mul(req.Amount, big.NewFloat(100)).Text('f', 0)

	transactionType := "STANDALONE_CREDIT"
	if int(Deposit) == req.OutType {
		transactionType = "PAYMENT"
	}

	requestData := map[string]interface{}{
		"merchantRefNum":  req.MerchantRefNum,
		"transactionType": transactionType,
		"paymentType":     "NETELLER",
		"amount":          amountInCents,
		"currencyCode":    req.Currency,
		"neteller": map[string]string{
			"consumerId": req.Email,
		},
		"returnLinks": []map[string]string{
			{"rel": "default", "href": "https://usgaminggamblig.com/payment/return/success"},
			{"rel": "on_failed", "href": "https://usgaminggamblig.com/payment/return/failed"},
			{"rel": "on_cancelled", "href": "https://usgaminggamblig.com/payment/return/cancel"},
		},
	}

	// Prepare headers
	authStr := base64.StdEncoding.EncodeToString([]byte(cli.MerchantID + ":" + cli.MerchantKey))
	encodedAuth := url.QueryEscape(authStr)

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Basic " + encodedAuth,
		"Simulator":     "EXTERNAL",
	}

	// Marshal request to JSON
	reqBody, err := json.Marshal(requestData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %v", err)
	}

	log.Printf("nettler#createPaymentHandle#req:%s", string(reqBody))

	// Create HTTP request
	httpReq, err := http.NewRequest("POST", cli.CreateHandleURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Set headers
	for k, v := range headers {
		httpReq.Header.Set(k, v)
	}

	// Send request
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Read response
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	log.Printf("nettler#createPaymentHandle#rsp:%s", string(respBody))

	// Parse response
	var response NetellerPaymentRsp
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	return &response, nil
}

// ------TODO 要再看一下这个process是做什么的?-----

func (cli *Client) Process(processReq *NetellerProcessReq) (*NetellerProcessRsp, error) {
	// Prepare request body
	requestMap := map[string]interface{}{
		"merchantRefNum":     processReq.MerchantRefNum,
		"amount":             cast.ToString(processReq.Amount), // big.Float to string
		"currencyCode":       processReq.CurrencyCode,
		"paymentHandleToken": processReq.PaymentHandleToken,
		"paymentType":        "NETELLER",
	}

	// Determine URL based on outType
	var requestURL string
	switch processReq.OutType {
	case int(Deposit):
		requestURL = cli.DepositURL
	case int(Withdraw):
		requestURL = cli.WithdrawURL
	default:
		return nil, fmt.Errorf("invalid outType: %d", processReq.OutType)
	}

	// Prepare authorization header
	authStr := base64.StdEncoding.EncodeToString([]byte(cli.MerchantID + ":" + cli.MerchantKey))
	encodedAuthStr := url.QueryEscape(authStr)

	// Marshal request body
	reqBody, err := json.Marshal(requestMap)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %v", err)
	}
	log.Printf("nettler#process#req: %s", string(reqBody))

	// Create HTTP request
	req, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+encodedAuthStr)
	req.Header.Set("Simulator", "EXTERNAL")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Read response
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}
	log.Printf("nettler#process#rsp: %s", string(respBody))

	// Parse response
	var processRsp NetellerProcessRsp
	if err := json.Unmarshal(respBody, &processRsp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	return &processRsp, nil
}
