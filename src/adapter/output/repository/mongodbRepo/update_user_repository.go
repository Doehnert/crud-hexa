package mongodbrepo

import (
	"context"
	"os"

	"github.com/Doehnert/crud-hexa/src/adapter/output/converter"
	"github.com/Doehnert/crud-hexa/src/application/domain"
	"github.com/Doehnert/crud-hexa/src/configuration/logger"
	"github.com/Doehnert/crud-hexa/src/configuration/rest_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) UpdateUser(userId string, userDomain domain.UserDomain,
) *rest_errors.RestErr {
	logger.Info("Init updateUser repo")
	collection_name := os.Getenv(MONGODB_USER_DB)
	collection := ur.databaseConnection.Collection(collection_name)

	domainEntity := converter.ConvertDomainToEntity(userDomain)
	userIdHex, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return rest_errors.NewBadRequestError("Invalid user id")
	}

	filter := bson.D{{Key: "_id", Value: userIdHex}}
	update := bson.D{{Key: "$set", Value: domainEntity}}

	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Error(
			"Error trying to update user",
			err,
			zap.String("journey", "updateUser"),
		)
		return rest_errors.NewInternalServerError(err.Error())
	}
	return nil
}
