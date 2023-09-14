package controller

import (
	"github.com/Doehnert/crud-hexa/src/application/port/input"
	"github.com/gin-gonic/gin"
)

func NewUserController(
	service input.UserDomainService,
) UserControllerInterface {
	return &userController{
		userService: service,
	}
}

type UserControllerInterface interface {
	FindUserByID(c *gin.Context)
	FindUserByEmail(c *gin.Context)
	CreateUser(c *gin.Context)
	LoginUser(c *gin.Context)
	UpdateUser(c *gin.Context)
}

type userController struct {
	userService input.UserDomainService
}
