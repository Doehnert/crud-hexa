package repository

import (
	"github.com/Doehnert/crud-hexa/src/application/port/output"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGODB_USER_DB = "MONGODB_USER_DB"
)

type userRepository struct {
	databaseConnection *mongo.Database
}

func NewUserRepository(
	database *mongo.Database,
) output.UserPort {
	return &userRepository{
		database,
	}
}
