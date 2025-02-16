package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/davg/logger/internal/domain/response"
)

const url = "http://users:8000/users"

func FetchUsers() ([]response.UserMetadata, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("ошибка запроса: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("неожиданный статус ответа: %d", resp.StatusCode)
	}

	var users []response.UserMetadata
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return nil, fmt.Errorf("ошибка декодирования JSON: %w", err)
	}

	return users, nil
}
