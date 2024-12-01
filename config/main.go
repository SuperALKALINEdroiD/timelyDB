package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type DatabaseConfig struct {
	DBName      string `json:"db_name"`
	DBType      string `json:"db_type"`
	TimeSeries  bool   `json:"time_series"`
	LockEnabled bool   `json:"lock_enabled"`
}

type Config struct {
	Databases []DatabaseConfig `json:"databases"`
}

func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %v", err)
	}

	defer file.Close()

	var config Config
	
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file: %v", err)
	}

	return &config, nil
}

func (c *Config) Append(newConfig DatabaseConfig) error {
	for _, dbConfig := range c.Databases {
		
		if dbConfig.DBName == newConfig.DBName {
			return fmt.Errorf("database with name %s already exists", newConfig.DBName)
		}
	}

	c.Databases = append(c.Databases, newConfig)
	return nil
}

func (c *Config) SaveConfig(filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("Failed to create config file: %v", err)
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ") 
	err = encoder.Encode(c)
	
	if err != nil {
		return fmt.Errorf("Failed to save Config File, Error on Encoder: %v", err)
	}

	return nil
}
