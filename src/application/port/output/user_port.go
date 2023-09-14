package output

import (
	"github.com/Doehnert/crud-hexa/src/application/domain"
	"github.com/Doehnert/crud-hexa/src/configuration/rest_errors"
)

type UserPort interface {
	CreateUser(
		userDomain domain.UserDomain,
	) (*domain.UserDomain, *rest_errors.RestErr)

	FindUserByEmail(
		email string,
	) (*domain.UserDomain, *rest_errors.RestErr)

	FindUserByID(
		id string,
	) (*domain.UserDomain, *rest_errors.RestErr)

	FindUserByEmailAndPassword(
		email, password string,
	) (*domain.UserDomain, *rest_errors.RestErr)

	UpdateUser(
		userId string,
		userDomain domain.UserDomain,
	) *rest_errors.RestErr

	DeleteUser(
		userId string,
	) *rest_errors.RestErr
}
