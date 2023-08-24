package router

import (
	"github.com/flpcastro/opportunities-api-go/internal/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	// Init Handler
	handler.InitHandler()

	v1 := router.Group("/api/v1")

	{
		v1.GET("/opportunities", handler.ListOpportunitiesHandler)
		v1.GET("/opportunity", handler.ShowOpportunityHandler)
		v1.POST("/opportunity", handler.CreateOpportunityHandler)
		v1.PUT("/opportunity", handler.UpdateOpportunityHandler)
		v1.DELETE("/opportunity", handler.DeleteOpportunityHandler)
	}

}
