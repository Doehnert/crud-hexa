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

func (uc *userController) LoginUser(c *gin.Context) {
	logger.Info("Init LoginUser controller")
	var userRequest request.LoginUserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to marshal object", err)
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	// Convert Request to Domain
	userDomain := &domain.UserDomain{
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}

	domainResult, token, err := uc.userService.LoginUserService(userDomain)
	if err != nil {
		logger.Error("Error trying to call LoginUser service",
			err,
			zap.String("journey", "loginUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("user logged in successfully", zap.String("joourney", "loginUser"))

	c.Header("Authorization", token)
	c.JSON(http.StatusOK, converter.ConvertDomainToResponse(domainResult))
}
