package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	db, err = InitializeSqlite()

	if err != nil {
		return fmt.Errorf("sqlite initialization error: %v", err)
	}
	return nil
}

func GetSqlite() *gorm.DB {
	return db
}

func GetLogger(prefix string) *Logger {
	logger = NewLogger(prefix)

	return logger
}
