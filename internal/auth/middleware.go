package auth

import (
	"net/http"
	"strings"

	"iq-go/internal/utils"

	"github.com/gin-gonic/gin"
)

func RequireAuth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		// Check for token in cookie for web requests
		token, err := c.Cookie("token")
		if err != nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}
		authHeader = "Bearer " + token
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == "" {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Authorization token required")
		c.Abort()
		return
	}

	claims, err := utils.ValidateToken(tokenString)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid token")
		c.Abort()
		return
	}

	c.Set("user_id", claims.UserID)
	c.Set("email", claims.Email)
	c.Next()
}
