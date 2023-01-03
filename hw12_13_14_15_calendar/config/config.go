package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App  `yaml:"app"`
	HTTP `yaml:"http"`
	Log  `yaml:"logger"`
	PG   `yaml:"postgres"`
}

type App struct {
	Name       string `env-required:"true" yaml:"name"    env:"APP_NAME"`
	Version    string `env-required:"true" yaml:"version"    env:"APP_VERSION"`
	APIVersion string `env-required:"false" yaml:"apiVersion"    env:"API_VERSION"`
	Storage    string `env-required:"true" yaml:"storage"    env:"APP_STOR"`
}

type HTTP struct {
	Host string `env-required:"true" yaml:"host"    env:"HTTP_HOST"`
	Port int    `env-required:"true" yaml:"port"    env:"HTTP_PORT"`
}

type Log struct {
	Level string `env-required:"true" yaml:"logLevel"    env:"LOG_LEVEL"`
}

type PG struct {
	// PoolMax int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"` // configure pool size
	URL string `env:"PG_URL"`
}

func NewConfig(path string) (*Config, error) {
	config := &Config{}
	err := cleanenv.ReadConfig(path, config)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
