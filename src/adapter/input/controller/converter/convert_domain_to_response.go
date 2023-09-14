package converter

import (
	"github.com/Doehnert/crud-hexa/src/adapter/input/model/response"
	"github.com/Doehnert/crud-hexa/src/application/domain"
)

func ConvertDomainToResponse(userDomain *domain.UserDomain) response.UserResponse {
	return response.UserResponse{
		ID:    userDomain.Id,
		Email: userDomain.Email,
		Name:  userDomain.Name,
		Age:   userDomain.Age,
	}
}
