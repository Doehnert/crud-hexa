package converter

import (
	"github.com/Doehnert/crud-hexa/src/adapter/input/model/response"
	"github.com/Doehnert/crud-hexa/src/application/domain"
)

func ConvertDomainToResponse(userDomain domain.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:    userDomain.GetId(),
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
