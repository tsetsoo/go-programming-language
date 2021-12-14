package github

import "time"

const IssuesURL = "https://api.github.com/search/issues"

const MyRepoIssuesURL = "https://api.github.com/repos/tsetsoo/go-programming-language/issues"

type IssuesSearchResults struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int       `json:"issue_number"`
	HTMLURL   string    `json:"html_url"`
	Title     string    `json:"title"`
	State     string    `json:"state"`
	User      *User     `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	Body      string    `json:"body"`
}

type User struct {
	Login   string `json:"login"`
	HTMLURL string `json:"html_url"`
}
