package domain

import (
	"database/sql/driver"
	"errors"
	"strings"
)

type MultiString []string

func (s *MultiString) Scan(src interface{}) error {
	str, ok := src.(string)
	if !ok {
		return errors.New("failed to scan multistring field - source is not a string")
	}
	*s = strings.Split(str, ",")
	return nil
}

func (s MultiString) Value() (driver.Value, error) {
	if (s == nil) || (len(s) == 0) {
		return nil, nil
	}
	return strings.Join(s, ","), nil
}

type OfferDocumentUpdate struct {
	ID string `gorm:"primaryKey" json:"id"`

	DocumentID string `gorm:"primaryKey" json:"document_id"`
	RecordID   string `gorm:"primaryKey" json:"record_id"` // к кому привязывается
	Type       string `json:"type"`
	URL        string `json:"url"`
}

type OfferRecordUpdate struct {
	ID string `gorm:"primaryKey" json:"id"`

	RecordID string `gorm:"primaryKey" json:"record_id"`

	Name         string      `json:"name"`
	MiddleName   string      `json:"middle_name"`
	LastName     string      `json:"last_name"`
	BirthDate    string      `json:"birth_date"`
	BirthPlace   string      `json:"birth_place"`
	MilitaryRank string      `json:"military_rank"`
	Commissariat string      `json:"commissariat"`
	Awards       MultiString `gorm:"type:text" json:"awards"`
	DeathDate    string      `json:"death_date"`
	BurialPlace  string      `json:"burial_place"`
	Bio          string      `json:"bio"`
}
