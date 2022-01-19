package omdbapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

const omdbApiUrl = "http://omdbapi.com/"

type omdbApiResult struct {
	Poster string `json:"Poster"`
}

func FindPosterUrl(title string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("%s?apikey=%s&t=%s", omdbApiUrl, os.Getenv("omdbapikey"), url.QueryEscape(title)))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result omdbApiResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}
	return result.Poster, nil
}
