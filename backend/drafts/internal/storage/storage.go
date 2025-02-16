package storage

import (
	"github.com/davg/drafts/internal/config"
	"github.com/davg/drafts/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func New() *Storage {
	cfg := config.Config().Storage

	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&domain.OfferDocumentUpdate{},
		&domain.OfferRecordUpdate{},
	)

	return &Storage{db: db}
}
