package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	err = InitializeEnvironmentVariables()
	if err!= nil {
    return fmt.Errorf("error initializing dotenv: %v", err.Error())
  }

	// Initialize SQlite database
	db, err = InitializeSqlite()
	if err!= nil {
    return fmt.Errorf("error initializing SQlite database: %v", err.Error())
  }

  return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	// Initialize logger
	logger = NewLogger(p)
	return logger
}