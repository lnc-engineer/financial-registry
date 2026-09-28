package app

import "errors"

// Config contains configuration for the registry application.
type Config struct {
	HTTPAddr string
}

// DefaultConfig returns the default application configuration.
func DefaultConfig() Config {
	return Config{
		HTTPAddr: ":8080",
	}
}

// Validate checks that the application configuration is valid.
func (c Config) Validate() error {
	if c.HTTPAddr == "" {
		return errors.New("http address must not be empty")
	}

	return nil
}
