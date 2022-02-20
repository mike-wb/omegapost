// handlers.go

package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mike-wb/omegapost/internal/models"
)

func ShowIndexPage(c *gin.Context) {
	articles := models.GetAllArticles()

	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles,
	}, "index.html")
}

func GetArticle(c *gin.Context) {
	// Check if the article ID is valid
	if articleId, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Check if the article exists
		if article, err := models.GetArticleByID(articleId); err == nil {
			// Call the HTML method of the Context to render a template
			render(c, gin.H{
				"title":   "Home Page",
				"payload": article,
			}, "index.html")
		} else {
			// If the article is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		// If an invalid article ID is provided in the URL, aborth with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}
