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

func TestDomainUserService_LoginUserService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserPort(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_calling_repo_returns_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		userDomain := &domain.UserDomain{
			Email:    "test@test.com",
			Password: "test",
			Name:     "test",
			Age:      90,
			Id:       id,
		}
		userDomainMock := &domain.UserDomain{
			Email:    userDomain.Email,
			Password: userDomain.Password,
			Name:     userDomain.Name,
			Age:      userDomain.Age,
			Id:       userDomain.Id,
		}
		userDomainMock.EncryptPassword()
		repository.EXPECT().FindUserByEmailAndPassword(userDomain.Email, userDomainMock.Password).Return(nil, rest_errors.NewInternalServerError("Error trying to find user by email and password"))

		userDomainReturn, token, err := service.LoginUserService(*userDomain)

		assert.Nil(t, userDomainReturn)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "Error trying to find user by email and password")
		assert.Empty(t, token)
	})

	t.Run("when_calling_create_token_returns_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		userDomain := &domain.UserDomain{
			Email:    "test@test.com",
			Password: "test",
			Name:     "test",
			Age:      90,
			Id:       id,
		}
		userDomainMock := mocks.NewMockUserDomainInterface(ctrl)
		userDomainMock.EXPECT().GenerateToken().Return("", rest_errors.NewInternalServerError("Error trying to generate token"))
		repository.EXPECT().FindUserByEmailAndPassword(gomock.Any(), gomock.Any()).Return(userDomain, nil)

		userDomainReturn, token, err := service.LoginUserService(*userDomain)

		assert.Nil(t, userDomainReturn)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "Error trying to generate token")
		assert.Empty(t, token)

	})
}
