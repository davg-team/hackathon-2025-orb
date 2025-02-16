package service

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

type OSMResponse struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

func GetCoordinates(address string) (float64, float64, error) {
	baseURL := "https://nominatim.openstreetmap.org/search"
	query := url.Values{}
	query.Set("q", address)
	query.Set("format", "json")

	resp, err := http.Get(baseURL + "?" + query.Encode())
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	var result []OSMResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil || len(result) == 0 {
		return 0, 0, fmt.Errorf("не найдено")
	}

	lat, lon := result[0].Lat, result[0].Lon
	latFloat, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		return 0, 0, err
	}

	lonFloat, err := strconv.ParseFloat(lon, 64)
	if err != nil {
		return 0, 0, err
	}

	latFloat, lonFloat = toEPSG3857(latFloat, lonFloat)
	return latFloat, lonFloat, nil
}

func toEPSG3857(lat, lon float64) (x, y float64) {
	x = lon * 20037508.34 / 180
	y = math.Log(math.Tan((90+lat)*math.Pi/360)) * 20037508.34 / math.Pi
	return x, y
}
