package logger

import (
	"github.com/caarlos0/env/v11"
	"go.uber.org/zap"
)

type Config struct {
	Level zap.AtomicLevel `env:"LEVEL" envDefault:"info"`
}

func LoadConfig() (Config, error) {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
