package config

import (
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	// db, err := gorm.Open("sqlite3", "test.db")

	return nil
}

func GetLogger(prefix string) *Logger {
	logger = NewLogger(prefix)

	return logger
}
