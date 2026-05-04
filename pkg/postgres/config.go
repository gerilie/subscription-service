package postgres

import "github.com/caarlos0/env/v11"

// Config represents configuration required to connect to a PostgreSQL database.
//
// Values are loaded from environment variables using github.com/caarlos0/env.
// Required fields must be non-empty.
type Config struct {
	Host     string `env:"HOST"              envDefault:"postgres"`
	Port     string `env:"PORT"              envDefault:"5432"`
	User     string `env:"USER,notEmpty"`
	Password string `env:"PASSWORD,notEmpty"`
	DB       string `env:"DB,notEmpty"`
}

// LoadConfig parses environment variables and returns a PostgreSQL configuration.
//
// It uses github.com/caarlos0/env for parsing. If required environment variables
// are missing or invalid, LoadConfig returns an error.
func LoadConfig() (Config, error) {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
