package main

import (
	"fmt"
	"os"

	"example.com/GithubFavorites"
	"github.com/joho/godotenv"
)

func main() {
	// make a cli app that fetches the top 5 favorites in GitHub.
	fmt.Println("Hello Golang!")
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		fmt.Println("Error loading .env file")
	}
	api := GithubFavorites.Api{
		Addr: "https://api.github.com",
		Key:  GithubFavorites.GetEnv("GITHUB_PAT", ""),
	}

	fmt.Print(api)
}
