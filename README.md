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

Copyright (c) 2026

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
