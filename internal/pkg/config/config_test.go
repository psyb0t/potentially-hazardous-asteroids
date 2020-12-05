package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

var (
	tmpConfigFpath   = fmt.Sprintf("/tmp/%d-config.testing.yaml", time.Now().Unix())
	tmpConfigContent = `debug: true
listen_address: "127.0.0.1:8080"
nasa_api_key: "testkey"
`
)

func setUp() error {
	err := os.RemoveAll(tmpConfigFpath)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(tmpConfigFpath, []byte(tmpConfigContent), 0644)
	if err != nil {
		return err
	}

	return nil
}

func tearDown() error {
	return os.RemoveAll(tmpConfigContent)
}

func TestConfigLoadFileSuccess(t *testing.T) {
	err := setUp()
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer tearDown()

	config := NewConfig()
	err = config.LoadFromFile(tmpConfigFpath)
	if err != nil {
		t.Fatalf(err.Error())
	}

	expectedDebug := true
	expectedListenAddress := "127.0.0.1:8080"
	expectedNASAAPIKey := "testkey"

	if config.Debug != expectedDebug {
		t.Fatalf("config.Debug value of %t is not the expected %t", config.Debug, expectedDebug)
	}

	if config.ListenAddress != expectedListenAddress {
		t.Fatalf("config.ListenAddress value of %s is not the expected %s", config.ListenAddress, expectedListenAddress)
	}

	if config.NASAAPIKey != expectedNASAAPIKey {
		t.Fatalf("config.NASAAPIKey value of %s is not the expected %s", config.NASAAPIKey, expectedNASAAPIKey)
	}
}

func TestConfigLoadFromFileFailOnInexistentFile(t *testing.T) {
	config := NewConfig()
	err := config.LoadFromFile("/inexistent-test-file-192938191")
	if err == nil {
		t.Fatalf("config.LoadFromFile falsely returned nil err on inexistent file")
	}
}

func TestConfigValidateSuccess(t *testing.T) {
	err := setUp()
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer tearDown()

	config := NewConfig()
	err = config.LoadFromFile(tmpConfigFpath)
	if err != nil {
		t.Fatalf(err.Error())
	}

	err = config.Validate()
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestConfigValidateFailOnNASAAPIKeyNotSet(t *testing.T) {
	err := setUp()
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer tearDown()

	config := NewConfig()
	err = config.LoadFromFile(tmpConfigFpath)
	if err != nil {
		t.Fatalf(err.Error())
	}

	config.NASAAPIKey = ""

	err = config.Validate()
	if !errors.Is(err, ErrNASAAPIKeyNotSet) {
		if err == nil {
			t.Fatalf("nil error falsely returned when NASA API key not set")
		}

		t.Fatalf(err.Error())
	}
}
