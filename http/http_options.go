package http

import (
	"time"

	log "github.com/sirupsen/logrus"
)

type HttpOption func(*client)

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
