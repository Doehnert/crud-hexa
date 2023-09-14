package mysqlrepo

import (
	"strconv"

	"github.com/Doehnert/crud-hexa/src/configuration/logger"
	"github.com/Doehnert/crud-hexa/src/configuration/rest_errors"
	"go.uber.org/zap"
)

func (ur *userRepository) DeleteUser(userId string) *rest_errors.RestErr {
	logger.Info("init DleteUser repo",
		zap.String("userId", userId))

	deleteQuery := "DELETE FROM users WHERE id = ?"

	userIDInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		logger.Error("Error parsing userID", err)
		return rest_errors.NewBadRequestError("Invalid userID")
	}

	_, err = ur.db.Exec(deleteQuery, userIDInt)
	if err != nil {
		return rest_errors.NewInternalServerError(err.Error())
	}
	return nil
}
