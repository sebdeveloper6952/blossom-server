package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

type DbConfig struct {
	Path         string `yaml:"path"`
	MigrationDir string `yaml:"migrations_dir"`
}

type StorageConfig struct {
	BasePath string `yaml:"base_path"`
}

type Config struct {
	Db      DbConfig      `yaml:"db"`
	Storage StorageConfig `yaml:"storage"`

	LogLevel           string   `yaml:"log_level"`
	ApiAddr            string   `yaml:"api_addr"`
	CdnUrl             string   `yaml:"cdn_url"`
	WhitelistedPubkeys []string `yaml:"whitelisted_pubkeys"`
}

func NewConfig(path string) (*Config, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = yaml.Unmarshal(bytes, config)

	return config, err
}
