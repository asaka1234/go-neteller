package go_neteller

import (
	"crypto/tls"
	"fmt"
	"github.com/asaka1234/go-neteller/utils"
	jsoniter "github.com/json-iterator/go"
)

// https://developer.paysafe.com/en/neteller-api-1/#/operations/get-payment-handle-using-merchant-reference-number
func (cli *Client) GetPaymentHandle(merchantRefNum string) (*NetellerGetPaymentHandleResp, error) {

	rawURL := cli.Params.GetPaymentHandleUrl

	//签名
	encodedAuth := utils.Sign(cli.Params.MerchantId, cli.Params.MerchantKey)

	//----------------------
	var result NetellerGetPaymentHandleResp

	resp2, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetQueryParam("merchantRefNum", merchantRefNum).
		SetHeaders(getAuthHeaders(encodedAuth)).
		SetDebug(cli.debugMode).
		SetResult(&result).
		SetError(&result).
		Get(rawURL)

	restLog, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(utils.GetRestyLog(resp2))
	cli.logger.Infof("PSPResty#neteller#GetPaymentHandle->%s", string(restLog))

	if err != nil {
		return nil, err
	}

	if resp2.Error() != nil {
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("%v", resp2.Error())
	}

	return &result, nil
}
