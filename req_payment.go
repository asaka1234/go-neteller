package go_neteller

import (
	"crypto/tls"
	"fmt"
	"github.com/asaka1234/go-neteller/utils"
	"github.com/mitchellh/mapstructure"
)

// https://developer.paysafe.com/en/neteller-api-1/#/#execution-mode
// https://developer.paysafe.com/en/neteller-api-1/#/operations/create-a-payment-handle
func (cli *Client) CreatePaymentHandle(req NetellerPaymentHandleReq) (*NetellerPaymentHandleResp, error) {

	rawURL := cli.Params.CreatePaymentHandleUrl

	var param map[string]interface{}
	mapstructure.Decode(req, &param)

	//补充字段
	param["paymentType"] = "NETELLER"
	param["returnLinks"] = []map[string]string{
		{"rel": "default", "href": cli.Params.PaymentBackUrl}, //回调地址
	}

	//签名
	encodedAuth := utils.Sign(cli.Params.MerchantId, cli.Params.MerchantKey)

	//----------------------
	var result NetellerPaymentHandleResp

	resp1, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetBody(param).
		SetHeaders(getAuthHeaders(encodedAuth)).
		SetResult(&result).
		SetError(&result).
		Post(rawURL)

	fmt.Printf("result: %s\n", string(resp1.Body()))

	if err != nil {
		return nil, err
	}

	return &result, nil
}
