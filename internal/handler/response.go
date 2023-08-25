package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type OpportunityResponse struct {
	ID        uint      `json:"id"`
	Role      string    `json:"role"`
	Company   string    `json:"company"`
	Location  string    `json:"location"`
	Remote    bool      `json:"remote"`
	Link      string    `json:"link"`
	Salary    int64     `json:"salary"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
}

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation %s successful", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateOpportunityResponse struct {
	Message string              `json:"message"`
	Data    OpportunityResponse `json:"data"`
}

type DeleteOpportunityResponse struct {
	Message string              `json:"message"`
	Data    OpportunityResponse `json:"data"`
}

type ShowOpportunityResponse struct {
	Message string              `json:"message"`
	Data    OpportunityResponse `json:"data"`
}

type ListOpportunitiesResponse struct {
	Message string                `json:"message"`
	Data    []OpportunityResponse `json:"data"`
}

type UpdateOpportunityResponse struct {
	Message string              `json:"message"`
	Data    OpportunityResponse `json:"data"`
}
