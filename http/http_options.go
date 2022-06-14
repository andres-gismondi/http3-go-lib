package http

import (
	"time"

	log "github.com/sirupsen/logrus"
)

type HttpOption func(*client)

func BaseURL(url string) HttpOption {
	return func(c *client) {
		c.options.baseURL = url
	}
}

func Headers(headers map[string]string) HttpOption {
	return func(c *client) {
		c.options.headers = headers
	}
}

func Logger(log *log.Logger) HttpOption {
	return func(c *client) {
		c.options.logger = log
	}
}

func Timeout(time time.Duration) HttpOption {
	return func(c *client) {
		c.options.timeout = time
	}
}
