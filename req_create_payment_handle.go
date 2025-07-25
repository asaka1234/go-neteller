package go_neteller

import (
	"crypto/tls"
	"fmt"
	"github.com/asaka1234/go-neteller/utils"
	jsoniter "github.com/json-iterator/go"
	"github.com/mitchellh/mapstructure"
)

// https://developer.paysafe.com/en/neteller-api-1/#/#execution-mode
// https://developer.paysafe.com/en/neteller-api-1/#/operations/create-a-payment-handle
func (cli *Client) CreatePaymentHandle(transactionType int, req NetellerPaymentHandleReq) (*NetellerPaymentHandleResp, error) {

	rawURL := cli.Params.CreatePaymentHandleUrl

	var param map[string]interface{}
	mapstructure.Decode(req, &param)

	//补充字段
	param["paymentType"] = "NETELLER"
	param["returnLinks"] = []map[string]string{
		{"rel": "default", "href": cli.Params.PaymentFeBackUrl}, //前端回跳地址
	}

	if transactionType == 1 {
		//deposit
		param["transactionType"] = "PAYMENT"
	} else if transactionType == 2 {
		//withdraw
		param["transactionType"] = "STANDALONE_CREDIT"
	}

	//签名
	encodedAuth := utils.Sign(cli.Params.MerchantId, cli.Params.MerchantKey)

	//----------------------
	var result NetellerPaymentHandleResp

	resp2, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetBody(param).
		SetHeaders(getAuthHeaders(encodedAuth)).
		SetDebug(cli.debugMode).
		SetResult(&result).
		SetError(&result).
		Post(rawURL)

	restLog, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(utils.GetRestyLog(resp2))
	cli.logger.Infof("PSPResty#neteller#CreatePaymentHandle->%s", string(restLog))

	if err != nil {
		return nil, err
	}

	if result.Error.Message != "" {
		// body里error
		return &result, nil
	}

	if resp2.Error() != nil {
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("%v", resp2.Error())
	}

	return &result, nil
}
