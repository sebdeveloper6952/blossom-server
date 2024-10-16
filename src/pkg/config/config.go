package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type AccessControlRule struct {
	Action   string `yaml:"action"`
	Pubkey   string `yaml:"pubkey"`
	Resource string `yaml:"resource"`
}

type Config struct {
	DbPath             string              `yaml:"db_path"`
	LogLevel           string              `yaml:"log_level"`
	ApiAddr            string              `yaml:"api_addr"`
	CdnUrl             string              `yaml:"cdn_url"`
	AdminPubkey        string              `yaml:"admin_pubkey"`
	MaxUploadSizeBytes int                 `yaml:"max_upload_size_bytes"`
	AccessControlRules []AccessControlRule `yaml:"access_control_rules"`
	AllowedMimeTypes   []string            `yaml:"allowed_mime_types"`
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
