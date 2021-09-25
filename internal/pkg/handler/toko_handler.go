package handler

import "github.com/gin-gonic/gin"

type TokoHandler struct {
}

func NewTokoHandler() *TokoHandler {
	return &TokoHandler{}
}

func (TokoHandler) Index(c *gin.Context) {

}

func (TokoHandler) Store(c *gin.Context) {

}

func (TokoHandler) Update(c *gin.Context) {

}

func (TokoHandler) Destroy(c *gin.Context) {

}
