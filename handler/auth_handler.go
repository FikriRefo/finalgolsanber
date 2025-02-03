package handler

import (
	"net/http"
	"socmed/dto"
	"socmed/errorhandle"
	"socmed/helper"
	"socmed/service"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	service service.AuthService
}

func NewAuthHandler(s service.AuthService) *authHandler {
	return &authHandler{
		service: s,
	}
}

func (h *authHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		errorhandle.HandleError(c, &errorhandle.BadRequestError{Message: err.Error()})
		return
	}

	if err := h.service.Register(&req); err != nil {
		errorhandle.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Berhasil Registrasi",
	})

	c.JSON(http.StatusCreated, res)
}

func (h *authHandler) Login(c *gin.Context) {
	var login dto.LoginRequest

	err := c.ShouldBindJSON(&login)
	if err != nil {
		errorhandle.HandleError(c, &errorhandle.BadRequestError{Message: err.Error()})
		return
	}

	result, err := h.service.Login(&login)
	if err != nil {
		errorhandle.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Berhasil Login",
		Data:       result,
	})

	c.JSON(http.StatusOK, res)
}
