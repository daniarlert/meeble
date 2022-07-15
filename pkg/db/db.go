package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"meeble/pkg/models"
)

type LocalDB struct {
	path string
	*gorm.DB
}

func NewLocalDB(path string) (*LocalDB, error) {
	conn, err := setup(path, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	ldb := &LocalDB{path, conn}
	return ldb, nil
}

func setup(path string, config *gorm.Config) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(path), config)
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&models.Mood{}); err != nil {
		return nil, err
	}

	return db, nil
}
