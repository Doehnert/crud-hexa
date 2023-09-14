package converter

import (
	"github.com/Doehnert/crud-hexa/src/adapter/output/model/entity"
	"github.com/Doehnert/crud-hexa/src/application/domain"
)

func ConvertDomainToEntity(
	domain domain.UserDomain,
) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    domain.Email,
		Password: domain.Password,
		Name:     domain.Name,
		Age:      domain.Age,
	}
}
