package config

import (
	"testing"
)

func TestLoadJSONConfig(t *testing.T) {
	config, err := LoadJSONConfig("../config.json")
	if err != nil {
		t.Fatalf("Failed to load JSON config: %v", err)
	}
	if config.Database.User != "json_user" {
		t.Errorf("Expected json_user, got %s", config.Database.User)
	}
}

func TestLoadYAMLConfig(t *testing.T) {
	config, err := LoadYAMLConfig("../config.yaml")
	if err != nil {
		t.Fatalf("Failed to load YAML config: %v", err)
	}
	if config.Database.User != "yaml_user" {
		t.Errorf("Expected yaml_user, got %s", config.Database.User)
	}
}

func TestLoadENVConfig(t *testing.T) {
	config, err := LoadENVConfig()
	if err != nil {
		t.Fatalf("Failed to load ENV config: %v", err)
	}
	if config.Database.User != "env_user" {
		t.Errorf("Expected env_user, got %s", config.Database.User)
	}
}
