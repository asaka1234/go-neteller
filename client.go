package go_neteller

import (
	"github.com/asaka1234/go-neteller/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	MerchantID  string // merchantId
	MerchantKey string // accessKey

	DepositURL      string
	WithdrawURL     string
	CreateHandleURL string

	ryClient *resty.Client
	logger   utils.Logger
}

func NewClient(logger utils.Logger, merchantID string, merchantKey, depositURL, withdrawURL, createHandleURL string) *Client {
	return &Client{
		MerchantID:      merchantID,
		MerchantKey:     merchantKey,
		DepositURL:      depositURL,
		WithdrawURL:     withdrawURL,
		CreateHandleURL: createHandleURL,

		ryClient: resty.New(), //client实例
		logger:   logger,
	}
}
