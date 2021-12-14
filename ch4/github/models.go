package github

import "time"

const IssuesURL = "https://api.github.com/search/issues"

const MyRepoIssuesURL = "https://api.github.com/repos/tsetsoo/programming-pearls/issues"

type IssuesSearchResults struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
