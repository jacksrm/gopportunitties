package config

import (
	"os"

	"github.com/jacksrm/gopportunitties/internal/opening/entity"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSqlite() (*gorm.DB, error) {
	const dbPath string = "./db/main.db"

	logger = GetLogger("sqlite")

	createDbIfNotExist(dbPath, logger)

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errorf("sqlite initialization error: %v\n", err)
		return nil, err
	}

	err = db.AutoMigrate(&entity.Opening{})
	if err != nil {
		logger.Errorf("sqlite migration error: %v\n", err)
		return nil, err
	}

	return db, nil
}

func createDbIfNotExist(dbPath string, logger *Logger) error {
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("sqlite database not found, creating new one")
		err := os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("sqlite initialization error: %v\n", err)
			return err
		}
		file, err := os.Create(dbPath)

		if err != nil {
			logger.Errorf("sqlite initialization error: %v\n", err)
			return err
		}

		file.Close()
	}

	return nil
}
