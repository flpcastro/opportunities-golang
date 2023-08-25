package handler

import (
	"net/http"

	"github.com/flpcastro/opportunities-api-go/internal/entity"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List opportunities
// @Description List all job opportunities
// @Tags Opportunities
// @Accept json
// @Produce json
// @Success 200 {object} ListOpportunitiesResponse
// @Failure 500 {object} ErrorResponse
// @Router /opportunities [get]
func ListOpportunitiesHandler(ctx *gin.Context) {
	opportunities := []entity.Opportunity{}

	if err := db.Find(&opportunities).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing opportunities")
		return
	}

	sendSuccess(ctx, "list-opportunities", opportunities)
}
