package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

// Config is returned by value because configuration is loaded
// once at startup and treated as read-only.
// also the config is small just few copies of slice and string headers, not the underlying data.
type Config struct {
	Stream StreamConfig
	Brands []Brand
	Output OutputConfig
}

type StreamConfig struct {
	URL string
}

type Brand struct {
	Name    string
	Domains []string
}

type OutputConfig struct {
	Level string
}

func Load(path string) (Config, error) {
	var cfg Config

	_, err := toml.DecodeFile(path, &cfg)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return Config{}, fmt.Errorf("could not find config file: %w", err)
		}
		return Config{}, fmt.Errorf("could not load config file: %w", err)
	}

	return cfg, nil

}
