package go_neteller

import (
	"fmt"
	"testing"
)

func TestWithdraw(t *testing.T) {

	//构造client
	cli := NewClient(nil, &NetellerInitParams{MERCHANT_ID, MERCHANT_KEY, CreatePaymentHandleUrl, ProcessStandaloneCreditsUrl, ProcessPaymentsUrl, GetPaymentHandleUrl, CreatePaymentHandleFeBackUrl})

	//发请求
	resp, err := cli.Withdraw(GenWithdrawRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

func GenWithdrawRequestDemo() NetellerPaymentHandleReq {
	return NetellerPaymentHandleReq{
		CurrencyCode:   "USD",
		Amount:         50000,
		MerchantRefNum: "7898", //商户订单号
		Neteller: NetellerDetail{
			ConsumerId: "demo@gmail.com",
		},
	}
}
