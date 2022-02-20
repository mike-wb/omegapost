// routes.go

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mike-wb/omegapost/internal/handlers"
)

func initializeRoutes(router *gin.Engine) {

	// Handle the index route
	router.GET("/", handlers.ShowIndexPage)
	// Handle GET requests at /article/view/some_article_id
	router.GET("/article/view/:article_id", handlers.GetArticle)
}
