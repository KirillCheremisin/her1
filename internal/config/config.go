package config

type Config struct {
	DefaultConfigPath string
	BaseURL           string `toml:"bind_BaseURL"`
}

func NewConfig() *Config {
	return &Config{
		DefaultConfigPath: "internal/config/config.toml",
		BaseURL:           "default url",
	}
}
