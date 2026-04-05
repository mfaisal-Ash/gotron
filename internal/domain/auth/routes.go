package auth

import (
	"diotron/internal/middleware"
	"github.com/gin-gonic/gin"
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
	common.JSONSuccessResponse(c, resp)
}


