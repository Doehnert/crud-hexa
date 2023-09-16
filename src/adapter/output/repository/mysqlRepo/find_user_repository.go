package mysqlrepo

import (
	"database/sql"
	"fmt"

	"github.com/Doehnert/crud-hexa/src/adapter/output/converter"
	"github.com/Doehnert/crud-hexa/src/adapter/output/model/entity"
	"github.com/Doehnert/crud-hexa/src/application/domain"
	"github.com/Doehnert/crud-hexa/src/configuration/logger"
	"github.com/Doehnert/crud-hexa/src/configuration/rest_errors"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmail(
	email string,
) (domain.UserDomainInterface, *rest_errors.RestErr) {
	logger.Info("Init findUserByEmail repo",
		zap.String("journey", "findUserByEmail"))

	selectQuery := "SELECT id, email, password, name, age FROM users WHERE email = ?"

	row := ur.db.QueryRow(selectQuery, email)

	userEntity := &entity.UserEntityMysql{}

	if err := row.Scan(
		&userEntity.ID,
		&userEntity.Email,
		&userEntity.Password,
		&userEntity.Name,
		&userEntity.Age,
	); err != nil {
		if err == sql.ErrNoRows {
			errorMessage := fmt.Sprint("User not found")
			return nil, rest_errors.NewNotFoundError(errorMessage)
		}
		return nil, rest_errors.NewInternalServerError(err.Error())
	}

	logger.Info("FindUserByEmail repo executed successfully",
		zap.String("journey", "findUserVyEmail"),
		zap.String("email", email),
	)
	return converter.ConvertMysqlEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByID(
	id string,
) (domain.UserDomainInterface, *rest_errors.RestErr) {
	logger.Info("Init findUseryID repo",
		zap.String("journey", "findUserByID"))

	selectQuery := "SELECT id, email, password, name, age FROM users WHERE id = ?"
	row := ur.db.QueryRow(selectQuery, id)

	userEntity := &entity.UserEntityMysql{}

	if err := row.Scan(
		&userEntity.ID,
		&userEntity.Email,
		&userEntity.Password,
		&userEntity.Name,
		&userEntity.Age,
	); err != nil {
		if err == sql.ErrNoRows {
			errorMessage := fmt.Sprint("User not found")
			return nil, rest_errors.NewNotFoundError(errorMessage)
		}
		return nil, rest_errors.NewInternalServerError(err.Error())
	}

	logger.Info("FindUserById repo executed successfully",
		zap.String("journey", "findUserByID"),
	)
	return converter.ConvertMysqlEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByEmailAndPassword(email, password string) (
	domain.UserDomainInterface, *rest_errors.RestErr) {
	logger.Info("Init findUser repo", zap.String("journey", "findUserByEmailAndPassword"))

	selectQuery := "SELECT id, email, password, name, age FROM users WHERE email = ? AND password = ?"
	row := ur.db.QueryRow(selectQuery, email, password)

	userEntity := &entity.UserEntityMysql{}

	if err := row.Scan(
		&userEntity.ID,
		&userEntity.Email,
		&userEntity.Password,
		&userEntity.Name,
		&userEntity.Age,
	); err != nil {
		if err == sql.ErrNoRows {
			errMessage := fmt.Sprint("user not found")
			return nil, rest_errors.NewNotFoundError(errMessage)
		}
		return nil, rest_errors.NewInternalServerError(err.Error())
	}

	logger.Info("FindUserByEmailAndPassword repo executed successfully",
		zap.String("journey", "findUserByEmailAndPassword"))
	zap.String("journey", email)
	return converter.ConvertMysqlEntityToDomain(*userEntity), nil
}
