package controller

import (
	"net/http"

	"github.com/Doehnert/crud-hexa/src/adapter/input/model/request"
	"github.com/Doehnert/crud-hexa/src/application/domain"
	"github.com/Doehnert/crud-hexa/src/configuration/logger"
	"github.com/Doehnert/crud-hexa/src/configuration/validation"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userController) UpdateUser(c *gin.Context) {
	logger.Info("Init UpdateUser controller",
		zap.String("journey", "update"))
	var userRequest request.UpdateUserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to marshal object", err)
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	userId := c.Param("userId")
	// if _, err := primitive.ObjectIDFromHex(userId); err != nil {
	// 	errRest := rest_errors.NewBadRequestError("Invalid userId, must be a hex value")
	// 	c.JSON(errRest.Code, errRest)
	// 	return
	// }

	// Convert Request to Domain
	userDomain := &domain.UserDomain{
		Name: userRequest.Name,
		Age:  userRequest.Age,
	}
	err := uc.userService.UpdateUserService(userId, userDomain)
	if err != nil {
		logger.Error("Error trying to call UpdateUser service",
			err,
			zap.String("journey", "updateUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("UpdateUser controller executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"))

	c.Status(http.StatusOK)
}
