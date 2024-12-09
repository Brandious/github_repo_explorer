package internals

import (
	"context"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/brandious/github_repo_explorer/views/pages"
	"github.com/gin-gonic/gin"
)

const appTimeout = time.Second * 10

func render(ctx *gin.Context, status int, template templ.Component) error {
	ctx.Header("Content-Type", "text/html; charset=utf-8")
	ctx.Status(status)

	return template.Render(ctx.Request.Context(), ctx.Writer)
}

func (app *Config) indexPageHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token, _ := ctx.Cookie("github_token")
		isAuthenticated := token != ""

		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		defer cancel()
		render(ctx, http.StatusOK, pages.Login(false, isAuthenticated))
	}
}

func (app *Config) reposPageHandler() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		token, _ := ctx.Cookie("github_token")
		isAuthenticated := token != ""

		_, cancel := context.WithTimeout(context.Background(), appTimeout)

		defer cancel()
		render(ctx, http.StatusOK, pages.Repos(true, isAuthenticated))
	}

}
