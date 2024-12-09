package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

func LoadEnv() {
	// Get the current working directory
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting working directory:", err)
	}

	// Load .env file
	err = godotenv.Load(filepath.Join(pwd, ".env"))
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}

func NewGithubOAuthConfig() *oauth2.Config {
	// Load environment variables from .env
	LoadEnv()

	clientID := os.Getenv("GITHUB_CLIENT_ID")
	clientSecret := os.Getenv("GITHUB_CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		log.Fatal("Missing GitHub OAuth credentials in .env file")
	}

	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{"user:email", "repo"},
		Endpoint:     github.Endpoint,
		RedirectURL:  "http://localhost:8080/callback",
	}
}
