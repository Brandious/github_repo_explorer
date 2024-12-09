package main

import (
	"fmt"

	"github.com/brandious/github_repo_explorer/cmd/internals"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/public", "./public")

	app := internals.Config{Router: router}
	app.Routes()

	fmt.Println("Server is running on port :8080")
	router.Run(":8080")
}
