package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type (
	Config struct {
		DB   DBConfig   `toml:"db"`
		HTTP HTTPConfig `toml:"http"`
	}

	DBConfig struct {
		Host     string `toml:"host"`
		Port     uint   `toml:"port"`
		DBName   string `toml:"dbname"`
		User     string `toml:"user"`
		Password string `toml:"password"`
	}

	HTTPConfig struct {
		Port uint `toml:"port"`
	}
)

func LoadConfig(path string) (*Config, error) {
	var cfg Config
	if _, err := toml.DecodeFile(path, &cfg); err != nil {
		return nil, fmt.Errorf("can not load config: %w", err)
	}
	return &cfg, nil
}
