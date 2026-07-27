package main

import (
	"fmt"
	"os"

	"example.com/GithubFavorites"
	"github.com/joho/godotenv"
)

func main() {
	// make a app that fetches the top 10 repos with the most stars made last week
	//  in GitHub.
	fmt.Println("Hello Golang!")
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		fmt.Println("Error loading .env file")
	}
	api := GithubFavorites.Api{
		Addr: GithubFavorites.GetEnv("GITHUB_URL", ""),
		Key:  GithubFavorites.GetEnv("GITHUB_PAT", ""),
	}
	repos, err := api.GetRepositories()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(repos)
}
