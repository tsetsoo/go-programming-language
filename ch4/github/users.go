package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ReadUsers() (*[]Issue, error) {
	resp, err := http.Get(MyRepoIssuesURL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result []Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
