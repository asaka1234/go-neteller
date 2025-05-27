package go_neteller

import (
	"github.com/asaka1234/go-neteller/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	MerchantID  string // merchantId
	MerchantKey string // accessKey

	CreatePaymentHandleURL string

	//回调地址
	PaymentBackURL string

	ryClient *resty.Client
	logger   utils.Logger
}

func NewClient(logger utils.Logger, merchantID string, merchantKey, createPaymentHandleURL, paymentBackURL string) *Client {
	return &Client{
		MerchantID:             merchantID,
		MerchantKey:            merchantKey,
		CreatePaymentHandleURL: createPaymentHandleURL,
		PaymentBackURL:         paymentBackURL,

		ryClient: resty.New(), //client实例
		logger:   logger,
	}
}
