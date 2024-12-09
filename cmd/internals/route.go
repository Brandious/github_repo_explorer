package internals

import (
	"github.com/brandious/github_repo_explorer/cmd/middleware"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Router *gin.Engine
}

func (app *Config) Routes() {

	app.Router.GET("/", app.indexPageHandler())
	app.Router.GET("/login", app.LoginHandler())
	app.Router.GET("/callback", app.CallbackHandler())

	protected := app.Router.Group("/")
	protected.Use(middleware.AuthRequired())
	{
		protected.GET("/repos", app.reposPageHandler())
		protected.GET("/search-repos", app.searchReposHandler())
	}
}
