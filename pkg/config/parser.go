package config

import (
	"github.com/BurntSushi/toml"
	"os"
	"path/filepath"
)

func Load() (*Config, error) {
	configPath := filepath.Join(os.Getenv("HOME"), ".config", "icebar", "config.toml")
	var cfg Config
	if _, err := toml.DecodeFile(configPath, &cfg); err != nil {
		return nil, err
	}

	appsPath := filepath.Join(os.Getenv("HOME"), ".config", "icebar", "apps.toml")
	var apps struct {
		Application []Application
	}
	if _, err := toml.DecodeFile(appsPath, &apps); err != nil {
		return nil, err
	}
	//make([]Application, len(apps.Application))
	cfg.Dock.Applications = apps.Application

	return &cfg, nil

}
