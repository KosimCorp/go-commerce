package handler

import (
	"ecommerce/internal/pkg/model"
	"ecommerce/internal/pkg/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) UserHandler {
	return UserHandler{userService: userService}
}

func (UserHandler) Index(c *gin.Context) {
	user := c.MustGet("user").(*model.User)

	c.JSON(200, *user)
}
