package go_neteller

import (
	"crypto/tls"
	"fmt"
	"github.com/asaka1234/go-neteller/utils"
	jsoniter "github.com/json-iterator/go"
	"github.com/mitchellh/mapstructure"
)

// https://developer.paysafe.com/en/neteller-api-1/#/operations/process-payments
func (cli *Client) ProcessPayments(req NetellerProcessPaymentsReq) (*NetellerProcessPaymentsResp, error) {

	rawURL := cli.Params.ProcessPaymentsUrl

	var param map[string]interface{}
	mapstructure.Decode(req, &param)

	//签名
	encodedAuth := utils.Sign(cli.Params.MerchantId, cli.Params.MerchantKey)

	//----------------------
	var result NetellerProcessPaymentsResp

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
	cli.logger.Infof("PSPResty#neteller#ProcessPayments->%s", string(restLog))

	if err != nil {
		return nil, err
	}

	if resp2.Error() != nil {
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("%v", resp2.Error())
	}

	return &result, nil
}
