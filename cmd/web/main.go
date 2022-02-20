// main.go

package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// Set the router as the default one provided by gin
	router = gin.Default()

	// Process the template at the start so they don't have to be loaded
	// from the disk again. This makes serving HTML pages fast.
	router.LoadHTMLGlob("templates/*")

	// Initialize the routes
	initializeRoutes(router)

	// Start serving the application
	router.Run()
}
