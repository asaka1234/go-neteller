package go_neteller

import (
	"fmt"
	"testing"
)

func TestDeposit(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &NetellerInitParams{MERCHANT_ID, MERCHANT_KEY, CreatePaymentHandleUrl, ProcessStandaloneCreditsUrl, ProcessPaymentsUrl, GetPaymentHandleUrl, CreatePaymentHandleFeBackUrl})

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
		Amount:         50000,
		MerchantRefNum: "21441113", //商户订单号
		Neteller: NetellerDetail{
			ConsumerId: "demo@gmail.com",
		},
	}
}
