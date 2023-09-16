package controller

import (
	"net/http"

	"github.com/Doehnert/crud-hexa/src/adapter/input/controller/converter"
	"github.com/Doehnert/crud-hexa/src/adapter/input/model/request"
	"github.com/Doehnert/crud-hexa/src/application/domain"
	"github.com/Doehnert/crud-hexa/src/configuration/logger"
	"github.com/Doehnert/crud-hexa/src/configuration/validation"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userController) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"))

	var userRequest request.CreateUserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info",
			err,
			zap.String("journey", "createUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}
	// Convert Request to Domain
	userDomain := domain.UserDomain{
		Email:    userRequest.Email,
		Password: userRequest.Password,
		Name:     userRequest.Name,
		Age:      userRequest.Age,
	}

	domainResult, err := uc.userService.CreateUserServices(&userDomain)
	if err != nil {
		logger.Error("Error trying to call CreateUser service",
			err,
			zap.String("journey", "createUser"))
		c.JSON(err.Code, err)
		return
	}
	logger.Info("CreateUser controller executed successfully",
		zap.String("userId", domainResult.GetId()),
		zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, converter.ConvertDomainToResponse(domainResult))
}
