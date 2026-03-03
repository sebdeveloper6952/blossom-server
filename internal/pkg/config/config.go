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

	config := &Config{
		DbPath:             "db/database.sqlite3",
		LogLevel:           "INFO",
		ApiAddr:            "0.0.0.0:8000",
		CdnUrl:             "http://0.0.0.0:8000",
		MaxUploadSizeBytes: 2097152, // 2MB
		AllowedMimeTypes:   []string{"*"},
	}
	err = yaml.Unmarshal(bytes, config)

	return config, err
}
