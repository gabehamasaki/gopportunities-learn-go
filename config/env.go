package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Enviroment struct {
	PORT string
}

func InitializeEnvironmentVariables() (*Enviroment, error) {
	
  if err := godotenv.Load(); err!= nil {
		logger.Errorf("error ")
    return nil, err
  }
	
	env := new(Enviroment);

	env.PORT = os.Getenv("PORT")

	return env, nil
}