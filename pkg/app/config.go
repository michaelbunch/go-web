package app

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Server configuration
type Config struct {
	AppName string       `yaml:"appName"`
	Server  ServerConfig `yaml:"server"`
}

// ServerConfig contains the configuration of the web server
type ServerConfig struct {
	Port int       `yaml:"port"`
	JWT  JWTConfig `yaml:"jwt"`
}

// JWTConfig contains JWT server settings
type JWTConfig struct {
	Issuer string `yaml:"issuer"`
	Secret string `yaml:"secret"`
}

// LoadConfig the config from the filesystem
func LoadConfig(configPath string) (*Config, error) {
	c := &Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)
	if err := d.Decode(&c); err != nil {
		return nil, err
	}

	return c, nil
}
