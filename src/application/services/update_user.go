package services

import (
	"github.com/Doehnert/crud-hexa/src/application/domain"
	"github.com/Doehnert/crud-hexa/src/configuration/logger"
	"github.com/Doehnert/crud-hexa/src/configuration/rest_errors"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUserService(userId string, userDomain domain.UserDomain) *rest_errors.RestErr {
	logger.Info("update user", zap.String("journey", "updateUser"))

	err := ud.repository.UpdateUser(userId, userDomain)
	if err != nil {
		logger.Error("Error trying to call UpdateUser repo",
			err,
			zap.String("journey", "updateUser"))
		return err
	}

	logger.Info("updateUser service executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"))

	return nil
}
