package middleware

import (
	"net/http"
	"strings"
	"diotron/internal/common"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware struct {
	JWTSecret string
}

func NewAuthMiddleware(secret string) *AuthMiddleware {
	return &AuthMiddleware{
		JWTSecret: secret,
	}
}

func (m *AuthMiddleware) MiddlewareFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			common.JSONErrorResponse(c, http.StatusUnauthorized, "Authorization header is required", nil)
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			common.JSONErrorResponse(c, http.StatusUnauthorized, "Authorization header format must be Bearer {token}", nil)
			c.Abort()
			return
		}

		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(m.JWTSecret), nil
		})

		if err != nil || !token.Valid {
			common.JSONErrorResponse(c, http.StatusUnauthorized, "Invalid token or expired token", nil)
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			common.JSONErrorResponse(c, http.StatusUnauthorized, "Invalid token claims", nil)
			c.Abort()
			return
		}

		UserID, ok := claims["user_id"].(string)
		if !ok {
			common.JSONErrorResponse(c, http.StatusUnauthorized, "Invalid token claims: user_id missing", nil)
			c.Abort()
			return
		}
		email, ok := claims["email"].(string)
		if !ok {
			common.JSONErrorResponse(c, http.StatusUnauthorized, "Invalid token claims: email missing", nil)
			c.Abort()
			return
		}
		role, ok := claims["role"].(string)
		if !ok {
			common.JSONErrorResponse(c, http.StatusUnauthorized, "Invalid token claims: role missing", nil)
			c.Abort()
			return
		}

		c.Set("user_id", UserID)
		c.Set("email", email)
		c.Set("role", role)
		c.Next()
	}

func (m *AuthMiddleware) requireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		roleVal, exists := c.Get("user_role")
		if !exists || roleVal != "admin" {
			common.JSONErrorResponse(c, http.StatusForbidden, "forbidden", nil)
			c.Abort()
			return
		}

		role, ok := roleVal.(string)
		if !ok || role != "admin" {
			common.JSONErrorResponse(c, http.StatusForbidden, "admin access required", nil)
			c.Abort()
			return
		}

		c.Next()
	}
}