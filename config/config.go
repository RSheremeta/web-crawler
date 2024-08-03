package config

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

const configPath = "./config.yaml"

type Config struct {
	Logger  *LoggerCfg  `yaml:"logger"`
	Printer *PrinterCfg `yaml:"printer"`
	Http    *HttpCfg    `yaml:"http"`

	DefaultTargetURL string `yaml:"default_target_url"`
}

type LoggerCfg struct {
	DebugLevel bool `yaml:"debug_level"`
}

type PrinterCfg struct {
	Throttling     time.Duration `yaml:"throttling"`
	ContextTimeout time.Duration `yaml:"context_timeout"`
}

type HttpCfg struct {
	Timeout             time.Duration `yaml:"timeout"`
	MaxIdleConns        int           `yaml:"max_idle_conns"`
	MaxIdleConnsPerHost int           `yaml:"max_idle_conns_per_host"`
	IdleConnTimeout     time.Duration `yaml:"idle_conn_timeout"`
}

func NewConfig() (*Config, error) {
	cfg, err := parseConfig()
	if err != nil {
		return nil, fmt.Errorf("config.parseConfig: %w", err)
	}

	return cfg, nil
}

func parseConfig() (*Config, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("os.ReadFile: %w", err)
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, fmt.Errorf("yaml.Unmarshal: %w", err)
	}

	return &cfg, nil
}
