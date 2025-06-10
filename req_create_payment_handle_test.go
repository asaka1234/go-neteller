package go_neteller

import (
	"fmt"
	"testing"
)

func TestDeposit(t *testing.T) {

	//构造client
	cli := NewClient(nil, &NetellerInitParams{MERCHANT_ID, MERCHANT_KEY, CreatePaymentHandleUrl, ProcessStandaloneCreditsUrl, ProcessPaymentsUrl, GetPaymentHandleUrl, CreatePaymentHandleFeBackUrl})

	//发请求
	resp, err := cli.Deposit(GenDepositRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

func GenDepositRequestDemo() NetellerPaymentHandleReq {
	return NetellerPaymentHandleReq{
		CurrencyCode:   "USD",
		CustomerIp:     "18.29.120.32",
		Amount:         50000,
		MerchantRefNum: "24567821113", //商户订单号
		Neteller: NetellerDetail{
			ConsumerId: "demo@gmail.com",
		},
	}
}
