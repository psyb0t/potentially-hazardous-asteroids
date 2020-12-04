package config

import "testing"

func TestConfigLoadFromFileFailOnInexistentFile(t *testing.T) {
	config := NewConfig()
	err := config.LoadFromFile("/inexistent-test-file-192938191")
	if err == nil {
		t.Fatalf("config.LoadFromFile falsely returned nil err on inexistent file")
	}
}
