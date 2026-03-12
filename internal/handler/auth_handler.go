package handler


import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/SauletTheBest/BackendFinancialApplication/internal/usecase"
)

type AuthHandler struct {
	authUsecase *usecase.AuthUsecase
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req usecase.RegisterRequest
}