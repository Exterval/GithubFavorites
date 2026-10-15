# GithubFavorites

A CLI tool that fetches the top 10 trending GitHub repositories from the past week, sorted by stars.

## Prerequisites

- [Go](https://go.dev/dl/) 1.26 or higher
- A [GitHub Personal Access Token](https://github.com/settings/tokens) (fine-grained, no scopes needed)

## Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/GithubFavorites.git
   cd GithubFavorites
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Create a `.env` file in the project root:
   ```env
   GITHUB_URL=https://api.github.com/search/repositories
   GITHUB_PAT=your_github_personal_access_token_here
   ```

## Usage

Run the CLI:
```bash
go run cmd/cli/main.go
```

## How It Works

1. Loads your GitHub PAT from the `.env` file
2. Queries the GitHub Search API for repos created in the last 7 days
3. Sorts results by stars (descending)
4. Returns the top 10 most popular new repositories

## License
MIT License
