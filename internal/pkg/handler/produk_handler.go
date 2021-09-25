package handler

import (
	"ecommerce/internal/pkg/service"

	"github.com/gin-gonic/gin"
)

type ProdukHandler struct {
	service *service.ProdukService
}

func NewProdukHandler(service *service.ProdukService) ProdukHandler {
	return ProdukHandler{service: service}
}

func (p *ProdukHandler) Index(c *gin.Context) {
	result, ok := p.service.FindAll()

	if !ok {
		c.AbortWithStatus(403)
		return
	}

	c.JSON(200, result)
}
