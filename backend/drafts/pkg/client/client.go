package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/davg/drafts/internal/domain"
	"io"
	"net/http"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
	token      string
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL:    baseURL,
		httpClient: &http.Client{},
	}
}

func (c *Client) SetToken(token string) {
	c.token = token
}

func (c *Client) UpdateRecord(id string, record *domain.OfferRecordUpdate) error {
	requestBody := struct {
		Name         string             `json:"name"`
		MiddleName   string             `json:"middle_name"`
		LastName     string             `json:"last_name"`
		BirthDate    string             `json:"birth_date"`
		BirthPlace   string             `json:"birth_place"`
		MilitaryRank string             `json:"military_rank"`
		Commissariat string             `json:"commissariat"`
		Awards       domain.MultiString `json:"awards"`
		DeathDate    string             `json:"death_date"`
		BurialPlace  string             `json:"burial_place"`
		Bio          string             `json:"bio"`
	}{
		Name:         record.Name,
		MiddleName:   record.MiddleName,
		LastName:     record.LastName,
		BirthDate:    record.BirthDate,
		BirthPlace:   record.BirthPlace,
		MilitaryRank: record.MilitaryRank,
		Commissariat: record.Commissariat,
		Awards:       record.Awards,
		DeathDate:    record.DeathDate,
		BurialPlace:  record.BurialPlace,
		Bio:          record.Bio,
	}

	url := fmt.Sprintf("%s/api/records/%s", c.baseURL, id)

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("error marshaling request: %w", err)
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(jsonData))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error executing request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, string(body))
	}

	return nil
}

func (c *Client) CreateDocument(recordID string, document *domain.OfferDocumentUpdate) error {
	if c.token == "" {
		return fmt.Errorf("authorization token is required")
	}

	requestBody := struct {
		Type string `json:"type"`
		URL  string `json:"url"`
	}{
		Type: document.Type,
		URL:  document.URL,
	}

	url := fmt.Sprintf("%s/api/records/%s/documents", c.baseURL, recordID)

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("error marshaling request: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonData))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error executing request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, string(body))
	}

	return nil
}
