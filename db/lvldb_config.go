package db

import "fmt"

type LevelDBConfig struct {
	Path string `json:"Path"`
}

func DefaultLevelDBConfig() LevelDBConfig {
	return LevelDBConfig{Path: "./data"}
}

func (cfg LevelDBConfig) IsValid() (bool, error) {
	if cfg.Path == "" {
		return false, fmt.Errorf("invalid path")
	}

	return true, nil
}
