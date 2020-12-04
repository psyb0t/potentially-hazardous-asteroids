package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/psyb0t/potentially-hazardous-asteroids/internal/app/pha"

	phaconfig "github.com/psyb0t/potentially-hazardous-asteroids/internal/pkg/config"
)

const (
	defaultListenAddress = "127.0.0.1:8080"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	log.SetLevel(log.WarnLevel)
	if _, isSet := os.LookupEnv("PHA_DEBUG"); isSet {
		log.SetLevel(log.DebugLevel)
	}
}

func main() {
	config, err := getConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	p := pha.NewPHA(config)

	log.Fatal(p.ListenAndServe())
}

// getConfig initializes a new pha config struct using either the file specified
// via the PHA_CONFIG_FPATH env var or other specific env vars to setup its fields
func getConfig() (*phaconfig.Config, error) {
	// Instantiate and load config struct
	config := phaconfig.NewConfig()

	// Get configuration file path from env variable
	configFilePath := os.Getenv("PHA_CONFIG_FPATH")
	if configFilePath != "" {
		err := config.LoadFromFile(configFilePath)
		if err != nil {
			return nil, err
		}
	}

	// If config's listen address is empty, try getting from the env var.
	// If not set, use default
	if config.ListenAddress == "" {
		config.ListenAddress = os.Getenv("PHA_LISTEN_ADDRESS")
		if config.ListenAddress == "" {
			log.Debugf("ListenAddress not specified! Using default ListenAddress %s", defaultListenAddress)

			config.ListenAddress = defaultListenAddress
		}
	}

	// If config's NASA API key is empty, try getting from the env var.
	// If not set, fail
	if config.NASAAPIKey == "" {
		config.NASAAPIKey = os.Getenv("PHA_NASA_API_KEY")
	}

	err := config.Validate()
	if err != nil {
		return nil, err
	}

	// If the PHA_DEBUG env var is set, config.Debug becomes true, whatever
	// the possibly set config file may say
	if _, isSet := os.LookupEnv("PHA_DEBUG"); isSet {
		config.Debug = true
	}

	return config, nil
}
