package config

import (
	"fmt"
	"strings"

	"github.com/marisasha/kinolog/internal/repository"
	"github.com/spf13/viper"
)

type Config struct {
	AppPort string            `mapstructure:"port"`
	DB      repository.Config `mapstructure:"db"`
}

func Load() (*Config, error) {

	if err := initConfig(); err != nil {
		return nil, fmt.Errorf("init config: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}

	return &cfg, nil
}

func initConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")

	viper.AutomaticEnv()

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	return viper.ReadInConfig()
}
