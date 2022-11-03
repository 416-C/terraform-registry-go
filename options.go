package terraformregistry

import "time"

type Option func(*Options)

type Options struct {
	BaseURL          string
	Timeout          time.Duration
	RetryCount       int
	RetryWaitTime    time.Duration
	RetryMaxWaitTime time.Duration
}

func WithTimeout(timeout time.Duration) Option {
	return func(o *Options) {
		o.Timeout = timeout
	}
}

func WithRetryCount(retryCount int) Option {
	return func(o *Options) {
		o.RetryCount = retryCount
	}
}

func WithRetryWaitTime(retryWaitTime time.Duration) Option {
	return func(o *Options) {
		o.RetryWaitTime = retryWaitTime
	}
}

func WithRetryMaxWaitTime(retryMaxWaitTime time.Duration) Option {
	return func(o *Options) {
		o.RetryMaxWaitTime = retryMaxWaitTime
	}
}
