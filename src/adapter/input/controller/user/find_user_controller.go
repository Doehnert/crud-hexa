package controller

import (
	"net/http"
	"net/mail"

	"github.com/Doehnert/crud-hexa/src/adapter/input/controller/converter"
	"github.com/Doehnert/crud-hexa/src/configuration/logger"
	"github.com/Doehnert/crud-hexa/src/configuration/rest_errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userController) FindUserByID(c *gin.Context) {
	logger.Info("Init findUserByID controller",
		zap.String("journey", "findUserByID"))

	userId := c.Param("userId")

	// Validate if userId is valid
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying to validate userId",
			err,
			zap.String("journey", "findUserById"))

		errorMessage := rest_errors.NewBadRequestError("UserID is not a valid id")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	domainResult, err := uc.userService.FindUserByIDServices(userId)
	if err != nil {
		logger.Error("Error trying to call findUserByID services",
			err,
			zap.String("journey", "findUserByID"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserByID controller executed successfully",
		zap.String("journey", "findUserByID"))
	c.JSON(http.StatusOK, converter.ConvertDomainToResponse(domainResult))
}

func (uc *userController) FindUserByEmail(c *gin.Context) {
	logger.Info("Init findUserByID controller",
		zap.String("journey", "findUserByID"))

	userEmail := c.Param("userEmail")

	if _, err := mail.ParseAddress(userEmail); err != nil {
		logger.Error("Error trying to validate userEmail",
			err,
			zap.String("journey", "findUserByEmail"))
		errorMessage := rest_errors.NewBadRequestError("userEmail is not a valid email")
		c.JSON(errorMessage.Code, errorMessage)
	}

	domainResult, err := uc.userService.FindUserByEmailServices(userEmail)
	if err != nil {
		logger.Error("Error trying to call findUserByEmail services",
			err,
			zap.String("journey", "findUserByEmail"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserByEmail controller executed successfully",
		zap.String("journey", "findUserByEmail"))
	c.JSON(http.StatusOK, converter.ConvertDomainToResponse(domainResult))
}
