package config

import (
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Database struct {
		User     string `json:"user" yaml:"user"`
		Password string `json:"password" yaml:"password"`
		DBName   string `json:"dbname" yaml:"dbname"`
	} `json:"database" yaml:"database"`
}

func LoadJSONConfig(filePath string) (*Configuration, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &Configuration{}
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func LoadYAMLConfig(filePath string) (*Configuration, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	config := &Configuration{}
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func LoadENVConfig() (*Configuration, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
		return nil, err
	}

	config := &Configuration{}
	config.Database.User = os.Getenv("DB_USER")
	config.Database.Password = os.Getenv("DB_PASSWORD")
	config.Database.DBName = os.Getenv("DB_NAME")
	return config, nil
}
