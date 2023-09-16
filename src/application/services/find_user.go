package services

import (
	"github.com/Doehnert/crud-hexa/src/application/domain"
	"github.com/Doehnert/crud-hexa/src/configuration/logger"
	"github.com/Doehnert/crud-hexa/src/configuration/rest_errors"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByIDServices(id string) (
	domain.UserDomainInterface, *rest_errors.RestErr,
) {
	logger.Info("Init findUserByID service",
		zap.String("journey", "findUserByID"))

	return ud.repository.FindUserByID(id)
}

func (ud *userDomainService) FindUserByEmailServices(email string) (
	domain.UserDomainInterface, *rest_errors.RestErr,
) {
	logger.Info("Init findUserByEmail service",
		zap.String("journey", "findUserByEmail"))

	return ud.repository.FindUserByEmail(email)
}

func (ud *userDomainService) FindUserByEmailAndPasswordServices(email string, password string) (
	domain.UserDomainInterface, *rest_errors.RestErr) {
	logger.Info("find user", zap.String("journey", "findUserByEmail"))

	return ud.repository.FindUserByEmailAndPassword(email, password)
}
