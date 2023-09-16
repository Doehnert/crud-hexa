package output

import (
	"github.com/Doehnert/crud-hexa/src/application/domain"
	"github.com/Doehnert/crud-hexa/src/configuration/rest_errors"
)

type UserPort interface {
	CreateUser(
		userDomain domain.UserDomainInterface,
	) (domain.UserDomainInterface, *rest_errors.RestErr)

	FindUserByEmail(
		email string,
	) (domain.UserDomainInterface, *rest_errors.RestErr)

	FindUserByID(
		id string,
	) (domain.UserDomainInterface, *rest_errors.RestErr)

	FindUserByEmailAndPassword(
		email, password string,
	) (domain.UserDomainInterface, *rest_errors.RestErr)

	UpdateUser(
		userId string,
		userDomain domain.UserDomainInterface,
	) *rest_errors.RestErr

	DeleteUser(
		userId string,
	) *rest_errors.RestErr
}
