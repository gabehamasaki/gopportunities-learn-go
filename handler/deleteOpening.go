package handler

import (
	"fmt"
	"net/http"

	"github.com/gabehamasaki/gopportunities-learn-go/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1/
// @Summary Delete opening
// @Description Delete a new job opening
// @Tags Openings
// @Accept  json
// @Produce  json
// @Param id query string true "Opening identificaton"
// @Success 200 {object} OpeningResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [delete]
func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id== "" {
		sendError(ctx, http.StatusBadRequest, 
			errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}
	// Find Opening
	if err:= db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, 
			fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	// Delete Opening
	if err:= db.Delete(&opening).Error; err!= nil {
    sendError(ctx, http.StatusInternalServerError, 
      fmt.Sprintf("error deletign opening with id: %s", id))
    return
  }

	sendSuccess(ctx, "delete-opening", opening)
}