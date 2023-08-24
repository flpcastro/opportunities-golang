package handler

import (
	"net/http"

	"github.com/flpcastro/opportunities-api-go/internal/entity"
	"github.com/gin-gonic/gin"
)

func ListOpportunitiesHandler(ctx *gin.Context) {
	opportunities := []entity.Opportunity{}

	if err := db.Find(&opportunities).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing opportunities")
		return
	}

	sendSuccess(ctx, "list-opportunities", opportunities)
}
