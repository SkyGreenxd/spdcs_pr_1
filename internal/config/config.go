package config

import (
	"github.com/SkyGreenxd/spdcs_pr_1/pkg/e"
	"github.com/joho/godotenv"
)

func LoadEnv() error {
	if err := godotenv.Load(); err != nil {
		return e.Wrap("error loading .env file", err)
	}

	return nil
}
