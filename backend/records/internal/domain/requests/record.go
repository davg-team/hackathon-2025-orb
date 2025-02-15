package requests

type RecordRequest struct {
	Name         string            `json:"name" binding:"required"`
	MiddleName   string            `json:"middle_name"`
	LastName     string            `json:"last_name" binding:"required"`
	BirthDate    string            `json:"birth_date"`
	BirthPlace   string            `json:"birth_place"`
	MilitaryRank string            `json:"military_rank"`
	Commissariat string            `json:"commissariat"`
	Awards       []string          `json:"awards"`
	DeathDate    string            `json:"death_date"`
	BurialPlace  string            `json:"burial_place"`
	Bio          string            `json:"bio"`
	Documents    []DocumentRequest `json:"documents"`
	Conflicts    []string          `json:"conflict_id"`
}

type DocumentRequest struct {
	Type     string `json:"type"`
	URL      string `json:"url"`
	RecordID string `json:"-"`
}
