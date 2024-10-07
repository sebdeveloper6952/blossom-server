package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	DbPath      string `yaml:"db_path"`
	LogLevel    string `yaml:"log_level"`
	ApiAddr     string `yaml:"api_addr"`
	CdnUrl      string `yaml:"cdn_url"`
	AdminPubkey string `yaml:"admin_pubkey"`
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
