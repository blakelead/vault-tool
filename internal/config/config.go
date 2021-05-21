package config

import "github.com/spf13/viper"

// Config contains Vault configuration from config file, cli and env
type Config struct {
	Address     string
	Token       string
	TLSInsecure bool
}

func GetSourceConfig() *Config {
	return getConfig("source")
}

func GetDestinationConfig() *Config {
	return getConfig("destination")
}

func getConfig(t string) *Config {
	return &Config{
		Address:     viper.GetString(t + ".address"),
		Token:       viper.GetString(t + ".token"),
		TLSInsecure: viper.GetBool(t + ".insecure"),
	}
}
