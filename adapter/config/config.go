package config

import (
	"go-gin-sqlc-template/domain"

	"github.com/spf13/viper"
)

// ReadConfig read config from local file
func ReadConfig() (*domain.Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AddConfigPath(".") // set config path to current work dir
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := &domain.Config{}
	err := viper.Unmarshal(cfg)

	return cfg, err
}
