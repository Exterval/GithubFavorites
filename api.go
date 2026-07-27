package GithubFavorites

import (
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
		panic(err)
	}
	req.Header.Set("Authorization", "Bearer "+a.Key)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", a.Version)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	// Check if validation failed
	if resp.StatusCode == http.StatusUnprocessableEntity {
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("Validation Error Details: %s\n", string(body))
		panic("Validation failed")
	}
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
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
	q.Add("page", "1")
	u.RawQuery = q.Encode()
	return u.String()
}
