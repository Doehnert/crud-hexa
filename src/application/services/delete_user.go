package services

import (
	"github.com/Doehnert/crud-hexa/src/configuration/logger"
	"github.com/Doehnert/crud-hexa/src/configuration/rest_errors"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUserService(userId string) *rest_errors.RestErr {
	logger.Info("init DeleteUser service",
		zap.String("journey", "deleteUser"))

	err := ud.repository.DeleteUser(userId)
	if err != nil {
		logger.Error("Error trying to call deleteUser repo",
			err,
			zap.String("journey", "deleteUser"))
		return err
	}

	logger.Info("deleteUser service executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"))

	return nil
}
