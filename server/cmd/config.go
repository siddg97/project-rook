package main

import (
	"github.com/siddg97/project-rook/pkg/utils"
)

type ServerConfig struct {
	LogLevel string
	Env      string
	Port     string
}

func InitConfig() (*ServerConfig, error) {
	envLogLevel, err := utils.GetDefaultEnvVar("LOG_LEVEL", "debug")
	if err != nil {
		return nil, err
	}

	envPort, err := utils.GetDefaultEnvVar("PORT", "3000")
	if err != nil {
		return nil, err
	}

	env, err := utils.GetDefaultEnvVar("ENV", "local")
	if err != nil {
		return nil, err
	}

	return &ServerConfig{
		LogLevel: envLogLevel,
		Env:      env,
		Port:     envPort,
	}, nil
}
