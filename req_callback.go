package go_neteller

import (
	"log"
)

// https://developer.paysafe.com/en/neteller-api-1/#/#webhooks-events
// http://paysafegroup.github.io/neteller_rest_api_v1/#/introduction/technical-introduction/webhooks
func (cli *Client) PaymentCallback(req NetellerPaymentBackReq, processor func(NetellerPaymentBackReq) error) error {
	log.Printf("Neteller#back#req: %+v", req)
	//TODO  是不是带了一个auth head过来? 需要用这个来验签?
	//开始处理
	return processor(req)
}
