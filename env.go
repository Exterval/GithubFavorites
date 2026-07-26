package GithubFavorites

import (
	"os"
)

func GetEnv(key, defaultValue string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return val
}
