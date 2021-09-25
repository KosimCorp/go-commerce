package middleware

import (
	"ecommerce/internal/pkg/service"

	"github.com/gin-gonic/gin"
)

func AccessToko(s *service.TokoService) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func AccessProduk(s *service.ProdukService) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
