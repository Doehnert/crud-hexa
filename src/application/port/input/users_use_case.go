package input

import (
	"github.com/Doehnert/crud-hexa/src/application/domain"
	"github.com/Doehnert/crud-hexa/src/configuration/rest_errors"
)

type UserDomainService interface {
	CreateUserServices(domain.UserDomain) (
		*domain.UserDomain, *rest_errors.RestErr)

	FindUserByIDServices(id string) (
		*domain.UserDomain, *rest_errors.RestErr)

	FindUserByEmailServices(email string) (
		*domain.UserDomain, *rest_errors.RestErr)

	LoginUserService(userDomain domain.UserDomain) (
		*domain.UserDomain, string, *rest_errors.RestErr)

	FindUserByEmailAndPasswordServices(email string, password string) (
		*domain.UserDomain, *rest_errors.RestErr)

	UpdateUserService(userId string, userDomain domain.UserDomain) *rest_errors.RestErr

	DeleteUserService(userId string) *rest_errors.RestErr
}
