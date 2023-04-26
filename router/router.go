package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize the router
	router := gin.Default()

	// Initialize routes
	initializeRoutes(router)

	// Start server
	router.Run(":8080")
}