package go_neteller

import (
	"crypto/tls"
	"fmt"
	"github.com/asaka1234/go-neteller/utils"
	"github.com/mitchellh/mapstructure"
)

// https://developer.paysafe.com/en/neteller-api-1/#/operations/process-standalone-credits
func (cli *Client) ProcessStandaloneCredits(req NetellerProcessStandaloneCreditsReq) (*NetellerProcessStandaloneCreditsResp, error) {

	rawURL := cli.Params.ProcessStandaloneCreditsUrl

	var param map[string]interface{}
	mapstructure.Decode(req, &param)

	//签名
	encodedAuth := utils.Sign(cli.Params.MerchantId, cli.Params.MerchantKey)

	//----------------------
	var result NetellerProcessStandaloneCreditsResp

	resp1, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetBody(param).
		SetHeaders(getAuthHeaders(encodedAuth)).
		SetDebug(cli.debugMode).
		SetResult(&result).
		SetError(&result).
		Post(rawURL)

	fmt.Printf("result: %s\n", string(resp1.Body()))

	if err != nil {
		return nil, err
	}

	return &result, nil
}
