package openMeteo

import (
	"errors"
	"fmt"
)

type Config struct {
	Host    string
	Version string
}

func (c *Config) Validate() error {
	var errorList []error

	if c.Host == "" {
		errorList = append(errorList, fmt.Errorf("open meteo host is required"))
	}
	if c.Version == "" {
		errorList = append(errorList, fmt.Errorf("open meteo host is required"))
	}

	return errors.Join(errorList...)
}
