package mysqlrepo

import (
	"strconv"

	"github.com/Doehnert/crud-hexa/src/adapter/output/converter"
	"github.com/Doehnert/crud-hexa/src/application/domain"
	"github.com/Doehnert/crud-hexa/src/configuration/rest_errors"
)

func (ur *userRepository) CreateUser(userDomain domain.UserDomainInterface) (domain.UserDomainInterface, *rest_errors.RestErr) {
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

	resultDomain := domain.NewUserDomain(
		strconv.FormatInt(insertedId, 10),
		userDomain.GetEmail(),
		userDomain.GetPassword(),
		userDomain.GetName(),
		userDomain.GetAge(),
	)

	return resultDomain, nil
}
