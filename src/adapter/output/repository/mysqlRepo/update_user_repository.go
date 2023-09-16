package mysqlrepo

import (
	"github.com/Doehnert/crud-hexa/src/adapter/output/converter"
	"github.com/Doehnert/crud-hexa/src/application/domain"
	"github.com/Doehnert/crud-hexa/src/configuration/logger"
	"github.com/Doehnert/crud-hexa/src/configuration/rest_errors"
)

func (ur *userRepository) UpdateUser(userId string, userDomain domain.UserDomainInterface,
) *rest_errors.RestErr {
	logger.Info("Init updateUser repo")

	updateQuery := `
    UPDATE users
    set email = ?, password = ?, name = ?, age = ?
    WHERE id = ?
  `

	entity := converter.ConvertDomainToEntityMysql(userDomain)
	_, err := ur.db.Exec(updateQuery, entity.Email, entity.Password, entity.Name, entity.Age, userId)
	if err != nil {
		return rest_errors.NewInternalServerError(err.Error())
	}
	return nil
}
