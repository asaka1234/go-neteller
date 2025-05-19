package go_praxis

import (
	"crypto/tls"
	"github.com/asaka1234/go-praxis/utils"
	"time"
)

// 下单(充值/提现是同一个接口)
func (cli *Client) Withdraw(req PraxisWithdrawReq) (*PraxisWithdrawResp, error) {

	rawURL := cli.BaseURL

	//拿到签名的参数
	requestParams := cli.CreateWithdrawRequestParams(req)
	requestSignatureList := cli.getRequestSignatureList()
	gtAuthenticationHeader := utils.GetGtAuthenticationHeader(requestParams, requestSignatureList, cli.MerchantKey)

	//返回值会放到这里
	var result PraxisWithdrawResp

	_, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetBody(requestParams).
		SetHeaders(getAuthHeaders(gtAuthenticationHeader)).
		SetResult(&result).
		Post(rawURL)

	//fmt.Printf("accessToken: %+v\n", resp)

	if err != nil {
		return nil, err
	}

	return &result, err
}

func (cli *Client) CreateWithdrawRequestParams(req PraxisWithdrawReq) map[string]interface{} {
	params := make(map[string]interface{})

	params["merchant_id"] = cli.MerchantID
	params["application_key"] = cli.ApplicationKey
	params["intent"] = req.Intent
	params["currency"] = req.Currency
	params["amount"] = req.Amount
	params["cid"] = req.Cid
	params["locale"] = cli.ApiLocale // Assuming Locale is a package constant
	params["customer_token"] = req.CustomerToken
	params["customer_data"] = req.CustomerData
	params["payment_method"] = req.PaymentMethod
	params["gateway"] = req.Gateway
	params["validation_url"] = req.ValidationURL
	params["notification_url"] = req.NotificationURL
	params["return_url"] = req.ReturnURL
	params["order_id"] = req.OrderID
	params["version"] = cli.ApiVersion      // Assuming APIVersion is a package constant
	params["timestamp"] = time.Now().Unix() // Unix timestamp in seconds

	return params
}
