package middleware

import (
	"ecommerce/internal/pkg/service"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(userService *service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()

		if !ok {
			c.AbortWithStatus(401)
			return
		}

		user, ok := userService.Login(username, password)
		if !ok {
			c.AbortWithStatus(401)
			return
		}

		c.Set("user", user)
	}
}
