package models

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
	if s == nil || len(s) == 0 {
		return nil, nil
	}
	return strings.Join(s, ","), nil
}

type RecordModel struct {
	ID           string          `gorm:"primaryKey" json:"id"`
	UserID       string          `json:"user_id"`
	Name         string          `json:"name"`
	MiddleName   string          `json:"middle_name"`
	LastName     string          `json:"last_name"`
	BirthDate    string          `json:"birth_date"`
	BirthPlace   string          `json:"birth_place"`
	MilitaryRank string          `json:"military_rank"`
	Commissariat string          `json:"commissariat"`
	Awards       MultiString     `gorm:"type:text" json:"awards"`
	DeathDate    string          `json:"death_date"`
	BurialPlace  string          `json:"burial_place"`
	Bio          string          `json:"bio"`
	MapID        string          `json:"map_id"`
	Documents    []DocumentModel `gorm:"foreignKey:RecordID" json:"documents"`
	Conflicts    []ConflictModel `gorm:"many2many:record_conflict;" json:"conflicts"`
	Published    bool            `json:"published"`
}

type DocumentModel struct {
	ID string `gorm:"primaryKey" json:"id"`

	RecordID string `json:"record_id"`
	Type     string `json:"type"`
	URL      string `json:"url"`
}

type ConflictModel struct {
	ID    string `gorm:"primaryKey" json:"id"`
	Title string `json:"title"`
	Dates string `json:"dates"`

	Records []RecordModel `gorm:"many2many:record_conflict;" json:"records"`
}
