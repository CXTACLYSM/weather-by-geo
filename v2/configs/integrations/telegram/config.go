package telegram

import (
	"errors"
	"fmt"
)

type Config struct {
	Host      string
	Token     string
	IsWebhook bool
}

func (c *Config) Validate() error {
	var errorList []error

	if c.Host == "" {
		errorList = append(errorList, fmt.Errorf("telegram host is required"))
	}
	if c.Token == "" {
		errorList = append(errorList, fmt.Errorf("telegram token is required"))
	}

	return errors.Join(errorList...)
}
