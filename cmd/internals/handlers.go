package internals

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/brandious/github_repo_explorer/cmd/config"
	"github.com/brandious/github_repo_explorer/views/components"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v60/github"
	"golang.org/x/oauth2"
)

func (app *Config) LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		state := generateRandomState()
		gitConfig := config.NewGithubOAuthConfig()
		// Save state to session/cookie
		c.SetCookie("oauth_state", state, 3600, "/", "localhost", false, true)

		// Redirect to GitHub
		url := gitConfig.AuthCodeURL(state)

		c.Redirect(http.StatusTemporaryRedirect, url)
	}
}

func (app *Config) CallbackHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get state from cookie
		state, _ := c.Cookie("oauth_state")
		gitConfig := config.NewGithubOAuthConfig()

		// Verify state
		if state != c.Query("state") {
			c.String(http.StatusBadRequest, "Invalid state")
			return
		}

		// Exchange code for token
		token, err := gitConfig.Exchange(c, c.Query("code"))
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to exchange token")
			return
		}

		// Create OAuth client
		client := github.NewClient(gitConfig.Client(c, token))

		// Get user info
		user, _, err := client.Users.Get(c, "")
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to get user info")
			return
		}

		userJSON, _ := json.MarshalIndent(user, "", "")
		fmt.Printf("GitHub User: %s\n", string(userJSON))

		// Store user info in session
		c.SetCookie("github_token", token.AccessToken, 3600, "/", "localhost", false, true)

		// Redirect to home page or dashboard
		c.Redirect(http.StatusTemporaryRedirect, "/repos")
	}
}

// Helper function to generate random state
func generateRandomState() string {
	// Implement your random string generation here
	return "random-state" // Use a proper random string in production
}

func (app *Config) searchReposHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.Query("repo")
		if query == "" {
			c.String(http.StatusBadRequest, "Repository name is required")
			return
		}

		// Get GitHub token from cookie
		token, err := c.Cookie("github_token")
		if err != nil {
			c.String(http.StatusUnauthorized, "Not authenticated")
			return
		}

		// Create GitHub client
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc := oauth2.NewClient(c, ts)
		client := github.NewClient(tc)

		// Search options
		opts := &github.SearchOptions{
			Sort:  "stars",
			Order: "desc",
			ListOptions: github.ListOptions{
				PerPage: 10,
			},
		}

		// Perform search
		result, _, err := client.Search.Repositories(c, query, opts)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to search repositories")
			return
		}

		// Render results
		render(c, http.StatusOK, components.SearchResults(result.Repositories))
	}
}
