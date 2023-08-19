package router

import (
	"github.com/flpcastro/opportunities-api-go/internal/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	{
		v1.GET("/openings", handler.ListOpportunitiesHandler)
		v1.GET("/opening", handler.ShowOpportunityHandler)
		v1.POST("/opening", handler.CreateOpportunityHandler)
		v1.PUT("/opening", handler.UpdateOpportunityHandler)
		v1.DELETE("/opening", handler.DeleteOpportunityHandler)
	}

}
