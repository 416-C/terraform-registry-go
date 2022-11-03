package terraformregistry

import (
	"time"

	"github.com/416-C/terraform-registry-go/modules"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	Modules *modules.Modules
}

func NewClient(baseURL string, opts ...Option) *Client {
	o := &Options{
		BaseURL:          baseURL,
		Timeout:          5 * time.Second,
		RetryCount:       3,
		RetryWaitTime:    5 * time.Second,
		RetryMaxWaitTime: 10 * time.Second,
	}

	for _, opt := range opts {
		opt(o)
	}

	client := resty.New().
		SetBaseURL(o.BaseURL).
		SetTimeout(o.Timeout).
		SetRetryCount(o.RetryCount).
		SetRetryWaitTime(o.RetryWaitTime).
		SetRetryMaxWaitTime(o.RetryMaxWaitTime)

	return &Client{
		Modules: modules.NewModulesService(client),
	}
}
