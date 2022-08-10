package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Role contains assume role config
type Role struct {
	Username    *string `yaml:"username"`
	Environment *string `yaml:"environment"`
	RoleARN     *string `yaml:"role_arn"`
	Region      *string `yaml:"region"`
}

// Config has root config
type Config struct {
	MasterAccountID *string `yaml:"master_account_id"`
	Roles           map[string]Role
}

func (config *Config) buildConfig(filename string) error {
	buffer, err := os.ReadFile(filename)
	if err != nil {
		return nil
	}
	err = yaml.Unmarshal(buffer, config)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("config error")
	}
	return nil
}

func getConfig(filename string) (*Config, error) {
	config := &Config{}
	err := config.buildConfig(filename)
	if err != nil {
		fmt.Println("Unable to produce a config")
		return nil, err
	}
	return config, nil
}
