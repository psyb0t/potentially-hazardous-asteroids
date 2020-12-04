package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config is the struct which holds configuration variables
type Config struct {
	Debug         bool   `yaml:"debug"`
	ListenAddress string `yaml:"listen_address"`
	NASAAPIKey    string `yaml:"nasa_api_key"`
}

// NewConfig is the constructor function which returns a
// pointer to a newly instantiated Config struct
func NewConfig() *Config {
	return &Config{}
}

// LoadFromFile reads the file residing at filePath and loads it
// into the instantiated Config struct
func (c *Config) LoadFromFile(filePath string) error {
	rawConfig, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(rawConfig, c)
}

// Validate checks if the required fields are set
func (c Config) Validate() error {
	if c.NASAAPIKey == "" {
		return ErrNASAAPIKeyNotSet
	}

	return nil
}
