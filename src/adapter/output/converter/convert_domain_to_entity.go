package converter

import (
	"github.com/Doehnert/crud-hexa/src/adapter/output/model/entity"
	"github.com/Doehnert/crud-hexa/src/application/domain"
)

func ConvertDomainToEntity(
	domain domain.UserDomainInterface,
) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Name:     domain.GetName(),
		Age:      domain.GetAge(),
	}
}
func ConvertDomainToEntityMysql(
	domain domain.UserDomainInterface,
) *entity.UserEntityMysql {
	return &entity.UserEntityMysql{
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Name:     domain.GetName(),
		Age:      domain.GetAge(),
	}
}
