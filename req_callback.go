package go_neteller

import (
	"fmt"
	"log"
)

// https://developer.paysafe.com/en/neteller-api-1/#/#webhooks-events
// http://paysafegroup.github.io/neteller_rest_api_v1/#/introduction/technical-introduction/webhooks
func (cli *Client) PaymentCallback(req NetellerPaymentBackReq, processor func(NetellerPaymentBackReq) error) error {
	log.Printf("Neteller#back#req: %+v", req)

	if req.EventType == string(PaymentHandlePayable) {
		//用户完成了支付,可以下一步了.

		handleInfo, err := cli.GetPaymentHandle(req.Payload.MerchantRefNum)
		if err != nil {
			return err
		}
		if handleInfo.PaymentHandles[0].Status != "PAYABLE" {
			return fmt.Errorf("%s", "wrong status")
		}

		//step-2
		processReq := NetellerProcessPaymentsReq{
			Amount:             req.Payload.Amount,
			CurrencyCode:       req.Payload.CurrencyCode,
			MerchantRefNum:     req.Payload.MerchantRefNum,
			PaymentHandleToken: handleInfo.PaymentHandles[0].PaymentHandleToken, //token
		}

		_, err = cli.ProcessPayments(processReq)
		if err != nil {
			return err
		}
		return nil
	}

	//开始处理
	return processor(req)
}
