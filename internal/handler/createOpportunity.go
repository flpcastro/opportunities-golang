package handler

import (
	"net/http"

	"github.com/flpcastro/opportunities-api-go/internal/entity"
	"github.com/gin-gonic/gin"
)

func CreateOpportunityHandler(ctx *gin.Context) {
	request := CreateOpportunityRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opportunity := entity.Opportunity{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opportunity).Error; err != nil {
		logger.Errorf("error creating opportunity: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opportunity on database")
		return
	}

	sendSuccess(ctx, "create opportunity", opportunity)
}
