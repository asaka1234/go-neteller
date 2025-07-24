package go_neteller

import "fmt"

// 集成接口
// TODO withdraw不需要等待回调
func (cli *Client) Withdraw(req NetellerPaymentHandleReq) (*NetellerProcessStandaloneCreditsResp, error) {

	//step-1, 创建句柄
	result, err := cli.CreatePaymentHandle(2, req)
	if err != nil {
		return nil, err
	}

	if result.Error.Message != "" {
		resp := &NetellerProcessStandaloneCreditsResp{}
		resp.BodyError = result.Error
		return resp, nil
	}

	//step-2, 处理请求
	//---------------------------------
	// withdraw: 可以直接下一步 (不需要等用户交互) . 二阶段请求发送完毕后,其状态为pending,  还是需要等待webhook来完成的.
	// deposit: 在webhook中下一步 (需要等用户交互)
	if result.Status == "PAYABLE" {

		processReq := NetellerProcessStandaloneCreditsReq{
			Amount:             req.Amount,
			CurrencyCode:       req.CurrencyCode,
			MerchantRefNum:     req.MerchantRefNum,
			PaymentHandleToken: result.PaymentHandleToken, //token
		}

		processRsp, err := cli.ProcessStandaloneCredits(processReq)
		if err != nil {
			return nil, err
		}

		return processRsp, nil
	}

	return nil, fmt.Errorf("handle status is %s", result.Status)
}
