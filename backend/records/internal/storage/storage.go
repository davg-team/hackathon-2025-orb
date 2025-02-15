package storage

import (
	"github.com/davg/records/internal/config"
	"github.com/davg/records/internal/domain/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func New() *Storage {
	cfg := config.Config().Storage

	db, err := gorm.Open(postgres.Open(cfg.URL), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.RecordModel{}, &models.DocumentModel{}, &models.ConflictModel{})

	return &Storage{db: db}
}
