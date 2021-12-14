package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func ReadIssues() (*[]Issue, error) {
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

func CreateIssue(toCreate *Issue) error {
	bytesBody, err := json.Marshal(toCreate)
	if err != nil {
		return err
	}
	body := bytes.NewBuffer(bytesBody)
	resp, err := http.Post(MyRepoIssuesURL, "application/json", body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 201 {
		return fmt.Errorf("Create issue failed with status code: %d and status: %s", resp.StatusCode, resp.Status)
	}
	return nil
}

func updateIssue(issueNumber string, toUpdate *Issue) error {
	bytesBody, err := json.Marshal(toUpdate)
	if err != nil {
		return err
	}

	body := bytes.NewBuffer(bytesBody)
	req, err := http.NewRequest(http.MethodPut, MyRepoIssuesURL, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("Update issue failed with status code: %d and status: %s", resp.StatusCode, resp.Status)
	}
	return nil
}
