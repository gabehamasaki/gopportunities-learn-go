package handler

import (
	"net/http"

	"github.com/gabehamasaki/gopportunities-learn-go/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1/
// @Summary List all openings
// @Description List all job openings
// @Tags Openings
// @Produce  json
// @Success 200 {object} OpeningsResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}