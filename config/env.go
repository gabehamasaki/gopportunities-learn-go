package config

import (
	"os"
)

var (
	PORT string
)

func InitializeEnvironmentVariables() error {
	
	
	PORT = os.Getenv("PORT")

	return nil
}