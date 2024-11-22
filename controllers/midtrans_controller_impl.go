package controllers

import (
	"DimasPrasetyo/backend-next/helpers"
	"DimasPrasetyo/backend-next/models/web"
	"DimasPrasetyo/backend-next/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MidtransControllerImpl struct {
	MidtransService services.MidtransService
}

func NewMidtransControllerImpl(midtransService services.MidtransService) *MidtransControllerImpl {
	return &MidtransControllerImpl{
		MidtransService: midtransService,
	}
}

func (controller *MidtransControllerImpl) Create(c *gin.Context) {
	var request web.MidtransRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		helpers.PanicIfError(err)
	}

	midtransResponse := controller.MidtransService.Create(c, request)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   midtransResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}