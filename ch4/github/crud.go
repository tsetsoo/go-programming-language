package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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

func ReadIssue(issueNumber int) (*Issue, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%d", MyRepoIssuesURL, issueNumber))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result Issue
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
	req, err := http.NewRequest("POST", MyRepoIssuesURL, body)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Authorization", "token "+os.Getenv("tokena"))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 201 {
		return fmt.Errorf("Create issue failed with status code: %d and status: %s", resp.StatusCode, resp.Status)
	}
	return nil
}

func UpdateIssue(issueNumber int, toUpdate *Issue) error {
	bytesBody, err := json.Marshal(toUpdate)
	if err != nil {
		return err
	}

	body := bytes.NewBuffer(bytesBody)
	issueUrl := fmt.Sprintf("%s/%d", MyRepoIssuesURL, issueNumber)
	fmt.Println(issueUrl)
	req, err := http.NewRequest(http.MethodPatch, issueUrl, body)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Authorization", "token "+os.Getenv("tokena"))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		newStr := buf.String()

		fmt.Printf(newStr)
		return fmt.Errorf("Update issue failed with status code: %d and status: %s", resp.StatusCode, resp.Status)
	}
	return nil
}
