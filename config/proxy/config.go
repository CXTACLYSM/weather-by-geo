package proxy

import (
	"errors"
	"fmt"
)

type Config struct {
	Host string
	Port string
}

func (c *Config) Validate() error {
	var errorList []error

	if c.Host == "" {
		errorList = append(errorList, errors.New("host is required"))
	}
	if c.Port == "" {
		errorList = append(errorList, errors.New("port is required"))
	}

	return errors.Join(errorList...)
}

func (c *Config) Url(protocol string) string {
	return fmt.Sprintf("%s://%s:%s", protocol, c.Host, c.Port)
}
