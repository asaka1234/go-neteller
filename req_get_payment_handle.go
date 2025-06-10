package go_neteller

import (
	"crypto/tls"
	"fmt"
	"github.com/asaka1234/go-neteller/utils"
)

// https://developer.paysafe.com/en/neteller-api-1/#/operations/get-payment-handle-using-merchant-reference-number
func (cli *Client) GetPaymentHandle(merchantRefNum string) (*NetellerGetPaymentHandleResp, error) {

	rawURL := cli.Params.GetPaymentHandleUrl

	//签名
	encodedAuth := utils.Sign(cli.Params.MerchantId, cli.Params.MerchantKey)

	//----------------------
	var result NetellerGetPaymentHandleResp

	resp1, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetQueryParam("merchantRefNum", merchantRefNum).
		SetHeaders(getAuthHeaders(encodedAuth)).
		SetDebug(cli.debugMode).
		SetResult(&result).
		SetError(&result).
		Get(rawURL)

	fmt.Printf("result: %s\n", string(resp1.Body()))

	if err != nil {
		return nil, err
	}

	return &result, nil
}
