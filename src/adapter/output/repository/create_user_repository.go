package repository

import (
	"context"
	"os"

	"github.com/Doehnert/crud-hexa/src/adapter/output/converter"
	"github.com/Doehnert/crud-hexa/src/application/domain"
	"github.com/Doehnert/crud-hexa/src/configuration/logger"
	"github.com/Doehnert/crud-hexa/src/configuration/rest_errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) CreateUser(
	userDomain domain.UserDomain,
) (*domain.UserDomain, *rest_errors.RestErr) {
	logger.Info("Init createUser repo",
		zap.String("journey", "createUser"))

	collection_name := os.Getenv(MONGODB_USER_DB)
	collection := ur.databaseConnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error("Error trying to create user",
			err,
			zap.String("journet", "createUser"))
		return nil, rest_errors.NewInternalServerError((err.Error()))
	}

	value.ID = result.InsertedID.(primitive.ObjectID)

	logger.Info(
		"CreateUser repo executed successfully",
		zap.String("userId", value.ID.Hex()),
		zap.String("journey", "createUser"),
	)

	return converter.ConvertEntityToDomain(*value), nil
}
