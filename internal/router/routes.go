package router

import (
	docs "github.com/flpcastro/opportunities-api-go/api"
	"github.com/flpcastro/opportunities-api-go/internal/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	// Init Handler
	handler.InitHandler()

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)

	{
		v1.GET("/opportunities", handler.ListOpportunitiesHandler)
		v1.GET("/opportunity", handler.ShowOpportunityHandler)
		v1.POST("/opportunity", handler.CreateOpportunityHandler)
		v1.PUT("/opportunity", handler.UpdateOpportunityHandler)
		v1.DELETE("/opportunity", handler.DeleteOpportunityHandler)
	}

	// Init Swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
