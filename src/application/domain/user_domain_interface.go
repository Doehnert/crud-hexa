package domain

import "github.com/Doehnert/crud-hexa/src/configuration/rest_errors"

type UserDomainInterface interface {
	EncryptPassword()
	GenerateToken() (string, *rest_errors.RestErr)
	GetId() string
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8
}

func NewUserDomain(id string, email string, password string, name string, age int8) UserDomainInterface {
	return &UserDomain{
		Id:       id,
		Email:    email,
		Password: password,
		Name:     name,
		Age:      age,
	}
}
