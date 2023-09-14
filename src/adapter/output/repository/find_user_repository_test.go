package repository

import (
	"fmt"
	"os"
	"testing"

	"github.com/Doehnert/crud-hexa/src/adapter/output/model/entity"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepo_FindUserByEmail(t *testing.T) {
	databaseName := "user_database_test"
	collection_name := "user_collection_test"
	err := os.Setenv("MONGODB_USER_DB", collection_name)
	if err != nil {
		t.FailNow()
	}
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mtestDb.Close()

	mtestDb.Run("when_sending_a_valid_email_returns_success", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "test@test.com",
			Password: "test",
			Name:     "test",
			Age:      50,
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", databaseName, collection_name),
			mtest.FirstBatch,
			convertEntityToBson(userEntity),
		))
		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepository(databaseMock)

		// Act
		userDomain, err := repo.FindUserByEmail(userEntity.Email)

		// Assert
		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.Id, userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.Email, userEntity.Email)
		assert.EqualValues(t, userDomain.Name, userEntity.Name)
		assert.EqualValues(t, userDomain.Age, userEntity.Age)
		assert.EqualValues(t, userDomain.Password, userEntity.Password)

	})

	mtestDb.Run("returns_error_when_mongodb_returns_error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepository(databaseMock)

		// Act
		userDomain, err := repo.FindUserByEmail("test@error.com")

		// Assert
		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})

	mtestDb.Run("returns_no_document_found", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", databaseName, collection_name),
			mtest.FirstBatch,
		))
		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepository(databaseMock)

		// Act
		userDomain, err := repo.FindUserByEmail("test@nouser.com")

		// Assert
		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})
}

func convertEntityToBson(userEntity entity.UserEntity) bson.D {
	return bson.D{
		{Key: "_id", Value: userEntity.ID},
		{Key: "email", Value: userEntity.Email},
		{Key: "password", Value: userEntity.Password},
		{Key: "name", Value: userEntity.Name},
		{Key: "age", Value: userEntity.Age},
	}
}
