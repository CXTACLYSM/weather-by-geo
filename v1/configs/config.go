package configs

import (
	"errors"
	"fmt"

	"github.com/CXTACLYSM/weather-by-geo/configs/app"
	"github.com/CXTACLYSM/weather-by-geo/configs/integrations/openMeteo"
	"github.com/CXTACLYSM/weather-by-geo/configs/integrations/telegram"
	"github.com/spf13/viper"
)

type Config struct {
	App       app.Config
	Telegram  telegram.Config
	OpenMeteo openMeteo.Config
}

func Create() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			return nil, fmt.Errorf("configs reading error: %w", err)
		}
	}

	config := &Config{
		App: app.Config{
			Host: viper.GetString("APP_HOST"),
			Port: viper.GetString("APP_PORT"),
		},
		Telegram: telegram.Config{
			Token: viper.GetString("TELEGRAM_TOKEN"),
			Host:  viper.GetString("TELEGRAM_HOST"),
		},
		OpenMeteo: openMeteo.Config{
			Host:    viper.GetString("OPEN_METEO_HOST"),
			Version: viper.GetString("OPEN_METEO_VERSION"),
		},
	}

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return config, nil
}

func (c *Config) Validate() error {
	return errors.Join(
		c.App.Validate(),
		c.Telegram.Validate(),
		c.OpenMeteo.Validate(),
	)
}
