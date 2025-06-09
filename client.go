package go_neteller

import (
	"github.com/asaka1234/go-neteller/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	Params *NetellerInitParams

	ryClient *resty.Client
	logger   utils.Logger
}

func NewClient(logger utils.Logger, params *NetellerInitParams) *Client {
	return &Client{
		Params: params,

		ryClient: resty.New(), //client实例
		logger:   logger,
	}
}
