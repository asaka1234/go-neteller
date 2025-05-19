package go_praxis

import (
	"github.com/asaka1234/go-praxis/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	MerchantID     string // merchantId
	MerchantKey    string // accessKey
	ApplicationKey string // merchantSecret
	ApiVersion     string
	ApiLocale      string

	BaseURL string

	ryClient *resty.Client
	logger   utils.Logger
}

func NewClient(logger utils.Logger, merchantID string, merchantKey, applicationKey, apiVersion, apiLocale string, baseURL string) *Client {
	return &Client{
		MerchantID:     merchantID,
		MerchantKey:    merchantKey,
		ApplicationKey: applicationKey,
		ApiVersion:     apiVersion,
		ApiLocale:      apiLocale,
		BaseURL:        baseURL,

		ryClient: resty.New(), //client实例
		logger:   logger,
	}
}
