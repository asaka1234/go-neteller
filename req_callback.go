package go_neteller

import (
	"log"
)

// 充值/提现的回调处理(传入一个处理函数)
func (cli *Client) CashierCallback(req NetellerBackReq, processor func(NetellerBackReq) error) error {
	log.Printf("Neteller#back#req: %+v", req)

	//开始处理
	return processor(req)
}
