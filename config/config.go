package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hedhyw/jsoncjson"
	"github.com/pkg/errors"
)

type Config struct {
	DB DB `json:"db"`
}

type DB struct {
	URL          string `json:"url"`
	SchemaName   string `json:"schema_name"`
	MaxOpenConns int    `json:"max_open_conns"`
	MaxIdleConns int    `json:"max_idle_conns"`
}

func DefaultConfig() (cfg *Config) {
	return &Config{
		DB: DB{
			URL:          "postgres://mydb1:123456@localhost:5432/mydb1?sslmode=disable",
			SchemaName:   "europe",
			MaxOpenConns: 2,
			MaxIdleConns: 2,
		},
	}
}

func LoadConfig(path string) (*Config, error) {
	cfg := DefaultConfig()

	confPath := os.Getenv("CONFIG_PATH")
	if confPath != "" {
		path = confPath
	}
	f, err := os.Open(path)
	if err != nil {
		return cfg, fmt.Errorf("opening: %w", err)
	}
	defer func() {
		_ = f.Close()
	}()
	jsoncReader := jsoncjson.NewReader(f)
	if err = json.NewDecoder(jsoncReader).Decode(cfg); err != nil {
		return nil, errors.WithStack(err)
	}
	return cfg, nil
}
