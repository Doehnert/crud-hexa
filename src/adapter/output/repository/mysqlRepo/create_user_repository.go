package mysqlrepo

import (
	"strconv"

	"github.com/Doehnert/crud-hexa/src/adapter/output/converter"
	"github.com/Doehnert/crud-hexa/src/application/domain"
	"github.com/Doehnert/crud-hexa/src/configuration/rest_errors"
)

func (ur *userRepository) CreateUser(userDomain domain.UserDomain) (*domain.UserDomain, *rest_errors.RestErr) {
	insertQuery := `
  INSERT INTO users (email, password, name, age)
  VALUES (?,?,?,?)
  `

	entity := converter.ConvertDomainToEntityMysql(userDomain)

	result, err := ur.db.Exec(insertQuery, entity.Email, entity.Password, entity.Name, entity.Age)
	if err != nil {
		return nil, rest_errors.NewInternalServerError(err.Error())
	}

	insertedId, _ := result.LastInsertId()

	resultDomain := &domain.UserDomain{
		Id:       strconv.FormatInt(insertedId, 10),
		Email:    userDomain.Email,
		Password: userDomain.Password,
		Name:     userDomain.Name,
		Age:      userDomain.Age,
	}

	return resultDomain, nil
}
