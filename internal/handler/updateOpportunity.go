package handler

import (
	"net/http"

	"github.com/flpcastro/opportunities-api-go/internal/entity"
	"github.com/gin-gonic/gin"
)

func UpdateOpportunityHandler(ctx *gin.Context) {
	request := UpdateOpportunityRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	opportunity := entity.Opportunity{}

	if err := db.First(&opportunity, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opportunity not found")
		return
	}

	if request.Role != "" {
		opportunity.Role = request.Role
	}

	if request.Company != "" {
		opportunity.Company = request.Company
	}

	if request.Location != "" {
		opportunity.Location = request.Location
	}

	if request.Remote != nil {
		opportunity.Remote = *request.Remote
	}

	if request.Link != "" {
		opportunity.Link = request.Link
	}

	if request.Salary > 0 {
		opportunity.Salary = request.Salary
	}

	if err := db.Save(&opportunity).Error; err != nil {
		logger.Errorf("error updating opportunity: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating opportunity")
		return
	}

	sendSuccess(ctx, "update-opportunity", opportunity)
}
