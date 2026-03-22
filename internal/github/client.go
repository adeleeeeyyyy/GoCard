package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"repo-card/internal/models"

)

func FetchRepo(owner, name string) (*models.GitHubRepo, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, name)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("repository not found (status: %d)", resp.StatusCode)
	}

	var repo models.GitHubRepo
	if err := json.NewDecoder(resp.Body).Decode(&repo); err != nil {
		return nil, err
	}
	return &repo, nil
}

func FetchReadme(owner, name, branch string) string {
	url := fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s/README.md", owner, name, branch)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return "readme doesn't provided"
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)
	text := string(data)
	if len(text) > 600 {
		return text[:600] + "\n\n... [Readme cutted]"
	}

	return text
}