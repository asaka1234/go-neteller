package go_neteller

// 下单(充值/提现是同一个接口)
func (cli *Client) Withdraw(req NetellerPaymentReq) (*NetellerPaymentRsp, error) {
	req.OutType = int(Withdraw)
	return cli.CreatePaymentHandle(&req)
}
