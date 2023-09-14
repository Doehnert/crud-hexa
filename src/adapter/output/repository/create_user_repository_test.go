package repository

import (
	"github.com/Doehnert/crud-hexa/src/application/domain"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"

	"os"
	"testing"
)

func TestUserRepo_CreateUser(t *testing.T) {
	databaseName := "user_database_test"
	collection_name := "user_collection_test"
	err := os.Setenv("MONGODB_USER_DB", collection_name)
	if err != nil {
		t.FailNow()
	}
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mtestDb.Close()

	mtestDb.Run("when_sending_a_valid_domain_returns_success", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})
		databaseMock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databaseMock)
		domain := domain.UserDomain{
			Email:    "test@test.com",
			Name:     "test",
			Password: "test",
			Age:      90,
		}

		// Act create user
		userDomain, err := repo.CreateUser(domain)

		// Asserts
		_, errId := primitive.ObjectIDFromHex(userDomain.Id)

		assert.Nil(t, err)
		assert.Nil(t, errId)
		assert.EqualValues(t, userDomain.Email, domain.Email)
		assert.EqualValues(t, userDomain.Name, domain.Name)
		assert.EqualValues(t, userDomain.Age, domain.Age)
		assert.EqualValues(t, userDomain.Password, domain.Password)

	})

	mtestDb.Run("return_error_from_database", func(mt *mtest.T) {
		// Arrange
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databaseMock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databaseMock)
		domain := domain.UserDomain{
			Email:    "test@test.com",
			Name:     "test",
			Password: "test",
			Age:      90,
		}

		// Act
		userDomain, err := repo.CreateUser(domain)

		//Assert
		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})
}
