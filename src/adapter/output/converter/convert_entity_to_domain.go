package converter

import (
	"strconv"

	"github.com/Doehnert/crud-hexa/src/adapter/output/model/entity"
	"github.com/Doehnert/crud-hexa/src/application/domain"
)

func ConvertEntityToDomain(
	entity entity.UserEntity,
) *domain.UserDomain {
	domainConverted := &domain.UserDomain{
		Email:    entity.Email,
		Password: entity.Password,
		Name:     entity.Name,
		Age:      entity.Age,
	}

	domainConverted.Id = entity.ID.Hex()
	return domainConverted
}

func ConvertMysqlEntityToDomain(
	entity entity.UserEntityMysql,
) domain.UserDomainInterface {
	domainConverted := &domain.UserDomain{
		Id:       strconv.FormatInt(entity.ID, 10),
		Email:    entity.Email,
		Password: entity.Password,
		Name:     entity.Name,
		Age:      entity.Age,
	}

	return domainConverted
}
