package auth

import (
	"gotron/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mfaisal-Ash/gotron/internal/common"
)

func RegisterRoutes(router *gin.Engine, handler *Handler, authMiddleware *middleware.AuthMiddleware) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", handler.Register)
		authGroup.POST("/login", handler.Login)
		authGroup.POST("/refresh", handler.RefreshToken)
	}
}

func RegisterProtectedRoutes(router *gin.Engine, handler *Handler, authMiddleware *middleware.AuthMiddleware) {
	protectedGroup := router.Group("/protected")
	protectedGroup.Use(authMiddleware.MiddlewareFunc())
	{
		protectedGroup.GET("/profile", handler.Profile)
	}
}

func (h *Handler) Profile(c *gin.Context) {
	claims, _ := c.Get("claims")
	userID := claims.(jwt.MapClaims)["user_id"].(string)
	resp, err := h.service.Me(userID)
	if err != nil {
		common.JSONErrorResponse(c, http.StatusNotFound, err.Error(), nil)
		return
	}
	common.JSONSuccessResponse(c, http.StatusOK, "Profile retrieved successfully", resp, nil)
}
