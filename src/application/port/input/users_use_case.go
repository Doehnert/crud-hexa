package input

import (
	"github.com/Doehnert/crud-hexa/src/application/domain"
	"github.com/Doehnert/crud-hexa/src/configuration/rest_errors"
)

type UserDomainService interface {
	CreateUserServices(domain.UserDomainInterface) (
		domain.UserDomainInterface, *rest_errors.RestErr)

	FindUserByIDServices(id string) (
		domain.UserDomainInterface, *rest_errors.RestErr)

	FindUserByEmailServices(email string) (
		domain.UserDomainInterface, *rest_errors.RestErr)

	LoginUserService(userDomain domain.UserDomainInterface) (
		domain.UserDomainInterface, string, *rest_errors.RestErr)

	FindUserByEmailAndPasswordServices(email string, password string) (
		domain.UserDomainInterface, *rest_errors.RestErr)

	UpdateUserService(userId string, userDomain domain.UserDomainInterface) *rest_errors.RestErr

	DeleteUserService(userId string) *rest_errors.RestErr
}
