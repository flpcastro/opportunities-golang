package handler

import (
	"fmt"
	"net/http"

	"github.com/flpcastro/opportunities-api-go/internal/entity"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Delete opportunity
// @Description Delete a new job opportunity
// @Tags Opportunities
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} DeleteOpportunityResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opportunity [delete]
func DeleteOpportunityHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opportunity := entity.Opportunity{}

	// Find Opening
	if err := db.First(&opportunity, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opportunity with id: %s not found", id))
		return
	}

	// Delete Opening
	if err := db.Delete(&opportunity).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opportunity with id: %s", id))
		return
	}

	sendSuccess(ctx, "delete-opportunity", opportunity)
}
