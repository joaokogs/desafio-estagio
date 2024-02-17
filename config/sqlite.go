package config

import (
	"os"

	"github.com/joaokogs/desafio-estagio/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitSQlite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating... ")

		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}
	// Migrate schemas
	err = db.AutoMigrate(&schemas.Missions{})
	if err != nil {
		logger.Errorf("SQlite automigration error: %v", err)
		return nil, err
	}

	return db, nil
}
