// render.go

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func render(c *gin.Context, data gin.H, templateName string) {
	// Look at the accept type to determine format to return the data
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}
}
