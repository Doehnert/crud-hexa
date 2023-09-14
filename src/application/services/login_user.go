package services

import (
	"github.com/Doehnert/crud-hexa/src/application/domain"
	"github.com/Doehnert/crud-hexa/src/configuration/logger"
	"github.com/Doehnert/crud-hexa/src/configuration/rest_errors"
	"go.uber.org/zap"
)

func (ud *userDomainService) LoginUserService(userDomain domain.UserDomain) (
	*domain.UserDomain, string, *rest_errors.RestErr) {
	logger.Info("Init login model", zap.String("journey", "loginUser"))

	userDomain.EncryptPassword()
	user, err := ud.FindUserByEmailAndPasswordServices(
		userDomain.Email,
		userDomain.Password,
	)

	if err != nil {
		return nil, "", err
	}

	token, err := user.GenerateToken()
	if err != nil {
		return nil, "", err
	}

	logger.Info("LoginUser service executed successfully",
		zap.String("userId", user.Id),
		zap.String("journey", "loginUser"))

	return user, token, nil
}
