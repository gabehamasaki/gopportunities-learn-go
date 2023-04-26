package router

import (
	"fmt"

	"github.com/gabehamasaki/gopportunities-learn-go/config"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	// Initialize the router
	router := gin.Default()

	// Initialize routes
	initializeRoutes(router)

	// Start server
	router.Run(fmt.Sprintf(":%s", config.Env.PORT))
}