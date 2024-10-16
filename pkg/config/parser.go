package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func Load() (*Config, error) {
	configPath := filepath.Join(os.Getenv("HOME"), ".config", "icebar", "config.toml")
	var cfg Config
	if _, err := toml.DecodeFile(configPath, &cfg); err != nil {
		return nil, err
	}

	appsPath := filepath.Join(os.Getenv("HOME"), ".config", "icebar", "apps.toml")
	var apps struct {
		Application []Application `toml:"applications"`
	}
	if _, err := toml.DecodeFile(appsPath, &apps); err != nil {
		return nil, err
	}
	//make([]Application, len(apps.Application))
	cfg.Dock.Applications = apps.Application
	log.Println("Box style form config: " + cfg.Appearance.BlockStyle + " path:" + cfg.Dock.Applications[0].Name)
	if !isValidHexColor(cfg.Appearance.MainColor) {
		return nil, fmt.Errorf("Invalid background color: %s", cfg.Appearance.MainColor)
	}

	if !isValidHexColor(cfg.Appearance.EdgeColor) {
		return nil, fmt.Errorf("Invalid edge color: %s", cfg.Appearance.EdgeColor)
	}

	return &cfg, nil

}

// Define a regex pattern for hex color codes
var hexColorRegex = regexp.MustCompile(`^#(?:[0-9a-fA-F]{3}){1,2}$`)

func isValidHexColor(s string) bool {
	return hexColorRegex.MatchString(s)
}
