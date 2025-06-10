package go_neteller

// 集成接口
func (cli *Client) Deposit(req NetellerPaymentHandleReq) (*NetellerPaymentHandleResp, error) {
	
	return cli.CreatePaymentHandle(1, req)
}
