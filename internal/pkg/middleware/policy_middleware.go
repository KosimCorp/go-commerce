package middleware

import (
	"ecommerce/internal/pkg/model"
	"ecommerce/internal/pkg/service"

	"github.com/gin-gonic/gin"
)

func AccessToko(s *service.TokoService) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("user").(*model.User)

		toko, ok := s.FindByUser(*user)

		if !ok {
			c.AbortWithStatus(403)
			return
		}

		c.Set("toko", toko)
	}
}
