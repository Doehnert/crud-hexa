package repository

import (
	"os"
	"testing"

	"github.com/Doehnert/crud-hexa/src/application/domain"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepo_UpdateUser(t *testing.T) {
	databaseName := "user_database_test"
	collection_name := "user_collection_test"
	err := os.Setenv("MONGODB_USER_DB", collection_name)
	if err != nil {
		t.FailNow()
		return
	}
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mtestDb.Close()

	mtestDb.Run("when_sending_valid_userId_and_userEntity_returns_success", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})
		userDomain := domain.UserDomain{
			Id:       primitive.NewObjectID().Hex(),
			Password: "test",
			Email:    "test@test.com",
			Name:     "test",
			Age:      59,
		}
		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepository(databaseMock)

		// Act
		err := repo.UpdateUser(userDomain.Id, userDomain)

		// Assert
		assert.Nil(t, err)
	})

	mtestDb.Run("return_error_when_database_return_error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		userDomain := domain.UserDomain{
			Id:       primitive.NewObjectID().Hex(),
			Password: "test",
			Email:    "test@test.com",
			Name:     "test",
			Age:      59,
		}
		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepository(databaseMock)

		// Act
		err := repo.UpdateUser(userDomain.Id, userDomain)

		// Assert
		assert.NotNil(t, err)
	})
}
