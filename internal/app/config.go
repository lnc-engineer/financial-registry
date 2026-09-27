package app

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
