package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/Doehnert/crud-hexa/src/adapter/output/converter"
	"github.com/Doehnert/crud-hexa/src/adapter/output/model/entity"
	"github.com/Doehnert/crud-hexa/src/application/domain"
	"github.com/Doehnert/crud-hexa/src/configuration/logger"
	"github.com/Doehnert/crud-hexa/src/configuration/rest_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmail(
	email string,
) (*domain.UserDomain, *rest_errors.RestErr) {
	logger.Info("Init findUserByEmail repo",
		zap.String("journey", "findUserByEmail"))

	collection_name := os.Getenv(MONGODB_USER_DB)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf(
				"User not found with this email: %s", email,
			)
			logger.Error(errorMessage, err, zap.String("journey", "findUserByEmail"))
			return nil, rest_errors.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by email"
		logger.Error(errorMessage, err, zap.String("journey", "findUserByEmail"))
		return nil, rest_errors.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserByEmail repo executed successfully",
		zap.String("journey", "findUserVyEmail"),
		zap.String("email", email),
		zap.String("userId", userEntity.ID.Hex()),
	)
	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByID(
	id string,
) (*domain.UserDomain, *rest_errors.RestErr) {
	logger.Info("Init findUseryID repo",
		zap.String("journey", "findUserByID"))

	collection_name := os.Getenv(MONGODB_USER_DB)

	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	objectId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id", Value: objectId}}
	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf(
				"User not found with this id: %s", id,
			)
			logger.Error(errorMessage, err, zap.String("journey", "findUserByID"))
			return nil, rest_errors.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by email"
		logger.Error(errorMessage, err, zap.String("journey", "findUserByID"))
		return nil, rest_errors.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserByEmail repo executed successfully",
		zap.String("journey", "findUserByID"),
		zap.String("userId", userEntity.ID.Hex()),
	)
	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByEmailAndPassword(email, password string) (
	*domain.UserDomain, *rest_errors.RestErr) {
	logger.Info("Init findUser repo", zap.String("journey", "findUserByEmailAndPassword"))

	collection_name := os.Getenv(MONGODB_USER_DB)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{
		{Key: "email", Value: email},
		{Key: "password", Value: password},
	}
	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(&userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprint("User or password is invalid")
			logger.Error(errorMessage, err, zap.String("journey", "findUserByEmail"))
			return nil, rest_errors.NewForbiddenError(errorMessage)
		}
	}

	logger.Info("FindUserByEmailAndPassword repo executed successfully",
		zap.String("journey", "findUserByEmailAndPassword"))
	zap.String("journey", email)
	zap.String("userId", userEntity.ID.Hex())
	return converter.ConvertEntityToDomain(*userEntity), nil
}
