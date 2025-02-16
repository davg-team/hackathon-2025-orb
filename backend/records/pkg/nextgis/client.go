package nextgis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

type Client struct {
	BaseURL  string
	LayerID  string
	Username string
	Password string
}

func NewClient(layerID, username, password string) *Client {
	return &Client{
		BaseURL:  "https://geois2.orb.ru/api",
		LayerID:  layerID,
		Username: username,
		Password: password,
	}
}

func (c *Client) doRequest(method, endpoint string, body io.Reader, contentType string) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s", c.BaseURL, endpoint)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.Username, c.Password)
	req.Header.Set("Accept", "*/*")
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	client := &http.Client{}
	return client.Do(req)
}

// Добавление записи
func (c *Client) AddFeature(data map[string]interface{}) (int, error) {
	url := fmt.Sprintf("resource/%s/feature/", c.LayerID)
	jsonData, _ := json.Marshal(data)
	resp, err := c.doRequest("POST", url, bytes.NewBuffer(jsonData), "application/json")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result map[string]int
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}
	return result["id"], nil
}

type Response struct {
	UploadMeta []UploadMetadata `json:"upload_meta"`
}

type UploadMetadata struct {
	ID       string `json:"id"`
	Size     int    `json:"size"`
	Name     string `json:"name"`
	MimeType string `json:"mime_type"`
}

// Загрузка файла
func (c *Client) UploadFile(fileBuffer *bytes.Buffer, fileName string) (*Response, error) {
	var b bytes.Buffer
	writer := multipart.NewWriter(&b)
	part, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, fileBuffer)
	if err != nil {
		return nil, err
	}
	writer.WriteField("name", fileName)
	writer.Close()

	resp, err := c.doRequest("POST", "component/file_upload/", &b, writer.FormDataContentType())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)

	var result *Response
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

// Прикрепление файла
func (c *Client) AttachFile(recordID int, fileID, fileName, mimeType string, size int) (int, error) {
	url := fmt.Sprintf("resource/%s/feature/%d/attachment/", c.LayerID, recordID)
	data := map[string]interface{}{
		"name":      fileName,
		"size":      size,
		"mime_type": mimeType,
		"file_upload": map[string]interface{}{
			"id":   fileID,
			"size": size,
		},
	}
	jsonData, _ := json.Marshal(data)
	resp, err := c.doRequest("POST", url, bytes.NewBuffer(jsonData), "application/json")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result map[string]int
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}
	return result["id"], nil
}

// Обновление записи
func (c *Client) UpdateFeature(recordID int, data map[string]interface{}) error {
	url := fmt.Sprintf("resource/%s/feature/%d", c.LayerID, recordID)
	jsonData, _ := json.Marshal(data)
	resp, err := c.doRequest("PUT", url, bytes.NewBuffer(jsonData), "application/json")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

// Удаление записи
func (c *Client) DeleteFeature(recordIDs []int) error {
	url := fmt.Sprintf("resource/%s/feature/", c.LayerID)
	jsonData, _ := json.Marshal(recordIDs)
	resp, err := c.doRequest("DELETE", url, bytes.NewBuffer(jsonData), "application/json")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

// Удаление вложения
func (c *Client) DeleteAttachment(recordID, attachmentID int) error {
	url := fmt.Sprintf("resource/%s/feature/%d/attachment/%d", c.LayerID, recordID, attachmentID)
	resp, err := c.doRequest("DELETE", url, nil, "")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

// Получение записей
func (c *Client) GetFeatures() ([]map[string]interface{}, error) {
	url := fmt.Sprintf("resource/%s/feature/", c.LayerID)
	resp, err := c.doRequest("GET", url, nil, "")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}
