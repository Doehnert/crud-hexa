package services

import (
	"github.com/Doehnert/crud-hexa/src/application/port/input"
	"github.com/Doehnert/crud-hexa/src/application/port/output"
)

func NewUserDomainService(repo output.UserPort) input.UserDomainService {
	return &userDomainService{
		repository: repo,
	}
}

type userDomainService struct {
	repository output.UserPort
}
