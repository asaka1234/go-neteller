package go_neteller

import (
	"fmt"
	"testing"
)

type VLog struct {
}

func (l VLog) Debugf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Infof(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Warnf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Errorf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

func TestWithdraw(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &NetellerInitParams{MERCHANT_ID, MERCHANT_KEY, CreatePaymentHandleUrl, ProcessStandaloneCreditsUrl, ProcessPaymentsUrl, GetPaymentHandleUrl, CreatePaymentHandleFeBackUrl})

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
