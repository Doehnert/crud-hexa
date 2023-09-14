package services

import (
	"testing"

	"github.com/Doehnert/crud-hexa/src/application/domain"
	"github.com/Doehnert/crud-hexa/src/configuration/rest_errors"
	"github.com/Doehnert/crud-hexa/src/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserDomainService_CreateUserServices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserPort(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_user_already_exists_returns_error", func(t *testing.T) {
		// Arrange
		id := primitive.NewObjectID().Hex()
		userDomain := &domain.UserDomain{
			Email:    "test@test.com",
			Password: "test",
			Name:     "test",
			Age:      90,
			Id:       id,
		}

		repository.EXPECT().FindUserByEmail(userDomain.Email).Return(
			userDomain, nil,
		)

		// Act
		user, err := service.CreateUserServices(*userDomain)

		// Assert
		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "Email is already registered in another account")

	})

	t.Run("when_user_is_not_registered_returns_error", func(t *testing.T) {
		// Arrange
		id := primitive.NewObjectID().Hex()
		userDomain := &domain.UserDomain{
			Email:    "test@test.com",
			Password: "test",
			Name:     "test",
			Age:      90,
			Id:       id,
		}

		repository.EXPECT().FindUserByEmail(userDomain.Email).Return(nil, nil)
		repository.EXPECT().CreateUser(gomock.Any()).Return(nil, rest_errors.NewInternalServerError("error trying to create user"))

		// Act
		user, err := service.CreateUserServices(*userDomain)

		// Assert
		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "error trying to create user")

	})
}
