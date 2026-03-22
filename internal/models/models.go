package models

type GithubOwner struct {
	Login     string `json:"login"`
	AvatarURL string `json:"avatar_url"`
	HTMLURL   string `json:"html_url"`
}

type GitHubRepo struct {
	Name            string      `json:"name"`
	FullName        string      `json:"full_name"`
	Description     string      `json:"description"`
	HTMLURL         string      `json:"html_url"`
	StargazersCount int         `json:"stargazers_count"`
	ForksCount      int         `json:"forks_count"`
	Language        string      `json:"language"`
	DefaultBranch   string      `json:"default_branch"`
	Owner           GithubOwner `json:"owner"`
}

type TemplateData struct {
	Repo   GitHubRepo
	Readme string
	Date   string
}
