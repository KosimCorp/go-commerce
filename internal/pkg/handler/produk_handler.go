package handler

import (
	"ecommerce/internal/pkg/model"
	"ecommerce/internal/pkg/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProdukHandler struct {
	service *service.ProdukService
}

func NewProdukHandler(service *service.ProdukService) ProdukHandler {
	return ProdukHandler{service: service}
}

func (p ProdukHandler) Index(c *gin.Context) {
	result, ok := p.service.FindAll()

	if !ok {
		c.AbortWithStatus(403)
		return
	}

	c.JSON(200, result)
}

func (p ProdukHandler) Show(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result, ok := p.service.Find(id)

	if !ok {
		c.AbortWithStatus(403)
		return
	}

	c.JSON(200, result)
}

func (p ProdukHandler) Store(c *gin.Context) {
	var produk model.Produk

	if err := c.ShouldBind(&produk); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ok := p.service.Create(&produk)

	if !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(200, produk)
}

func (p ProdukHandler) Update(c *gin.Context) {

}

func (p ProdukHandler) Destroy(c *gin.Context) {

}
