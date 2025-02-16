package requests

import "github.com/davg/drafts/internal/domain"

type UpdateDocumentPost struct {
	RecordID string `json:"record_id"`
	Type     string `json:"type"`
	URL      string `json:"url"`
}

type UpdateRecordPost struct {
	RecordID     string             `json:"record_id"`
	Name         string             `json:"name"`
	MiddleName   string             `json:"middle_name"`
	LastName     string             `json:"last_name"`
	BirthDate    string             `json:"birth_date"`
	BirthPlace   string             `json:"birth_place"`
	MilitaryRank string             `json:"military_rank"`
	Commissariat string             `json:"commissariat"`
	Awards       domain.MultiString `gorm:"type:text" json:"awards"`
	DeathDate    string             `json:"death_date"`
	BurialPlace  string             `json:"burial_place"`
	Bio          string             `json:"bio"`
}
