package services

import (
	"DimasPrasetyo/backend-next/models/web"
	"github.com/gin-gonic/gin"
)

type MidtransService interface {
	Create(c *gin.Context, request web.MidtransRequest) web.MidtransResponse
}