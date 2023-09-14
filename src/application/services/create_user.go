package services

import (
	"github.com/Doehnert/crud-hexa/src/application/domain"
	"github.com/Doehnert/crud-hexa/src/configuration/logger"
	"github.com/Doehnert/crud-hexa/src/configuration/rest_errors"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUserServices(userDomain domain.UserDomain) (
	*domain.UserDomain, *rest_errors.RestErr,
) {
	logger.Info("Init createUser service",
		zap.String("journey", "createUser"))

	user, _ := ud.FindUserByEmailServices(userDomain.Email)
	if user != nil {
		return nil, rest_errors.NewBadRequestError("Email is already registered in another account")
	}

	userDomain.EncryptPassword()
	userDomainRepo, err := ud.repository.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error trying to call repo",
			err,
			zap.String("journey", "createUser"))
		return nil, err
	}

	logger.Info("CreateUser service executed successfully",
		zap.String("userId", userDomainRepo.Id),
		zap.String("journey", "createUser"))

	return userDomainRepo, nil
}
