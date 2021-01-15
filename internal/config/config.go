package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// Environment is an enumerator to represent the environment the application is running in.
type Environment int

// Environment is an enumerator to represent the environment the application is running in.
const (
	Development Environment = iota
	Production
)

// Config a struct that contains all your configuerations
type Config struct {
	Environment Environment `mapstructure:"ENVIRONMENT" default:"1"`
}

// Load unmarshals configuerations into an instance of struct `Settings`
func Load() (*Config, error) {
	c := Config{}
	return &c, viper.Unmarshal(&c)
}

// InitConfig reads in config file and ENV variables if set.
func InitConfig(cfgFile string) {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	}

	viper.AutomaticEnv() // read in environment variables that match
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
