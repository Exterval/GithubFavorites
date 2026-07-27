package GithubFavorites

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Api struct {
	Addr    string
	Key     string
	Version string
}

type Repository struct {
	Name        string
	Description string
	URL         string
}

func (a *Api) GetRepositories() ([]Repository, error) {
	repos := []Repository{}
	url := a.makeUrl()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+a.Key)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", a.Version)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: %s - %s", resp.Status, string(body))
	}

	var result struct {
		Items []struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			HTMLURL     string `json:"html_url"`
		} `json:"items"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// Map top 10 repos to our struct
	for i, item := range result.Items {
		if i >= 10 {
			break
		}
		repos = append(repos, Repository{
			Name:        item.Name,
			Description: item.Description,
			URL:         item.HTMLURL,
		})
	}

	return repos, nil
}

func (a *Api) makeUrl() string {
	u, err := url.Parse(a.Addr)
	if err != nil {
		panic(err)
	}
	q := u.Query()
	now := time.Now()
	week := now.Add(-7 * 24 * time.Hour)
	weekFormatted := week.Format("2006-01-02")
	q.Set("q", "created:"+weekFormatted)
	q.Add("sort", "stars")
	q.Add("order", "desc")
	q.Add("per_page", "10")
	u.RawQuery = q.Encode()
	return u.String()
}
