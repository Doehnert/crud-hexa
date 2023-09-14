package services

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/Doehnert/crud-hexa/src/application/domain"
	"github.com/Doehnert/crud-hexa/src/configuration/rest_errors"
	"github.com/Doehnert/crud-hexa/src/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserDomainService_FindUserByIDServices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserPort(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_exists_an_user_returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		userDomain := &domain.UserDomain{
			Email:    "test@test.com",
			Password: "test",
			Name:     "test",
			Age:      90,
			Id:       id,
		}
		repository.EXPECT().FindUserByID(id).Return(userDomain, nil)

		userDomainReturn, err := service.FindUserByIDServices(id)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomainReturn.Id, id)
		assert.EqualValues(t, userDomainReturn.Email, userDomain.Email)
		assert.EqualValues(t, userDomainReturn.Password, userDomain.Password)
		assert.EqualValues(t, userDomainReturn.Name, userDomain.Name)
		assert.EqualValues(t, userDomainReturn.Age, userDomain.Age)
	})

	t.Run("when_does_not_exists_an_user_returns_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		repository.EXPECT().FindUserByID(id).Return(nil, rest_errors.NewNotFoundError("user not found"))

		userDomainReturn, err := service.FindUserByIDServices(id)

		assert.Nil(t, userDomainReturn)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "user not found")

	})
}
func TestUserDomainService_FindUserByEmailServices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserPort(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_exists_an_user_returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		email := "test@success.com"

		userDomain := &domain.UserDomain{
			Email:    email,
			Password: "test",
			Name:     "test",
			Age:      90,
			Id:       id,
		}
		repository.EXPECT().FindUserByEmail(email).Return(userDomain, nil)

		userDomainReturn, err := service.FindUserByEmailServices(email)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomainReturn.Id, id)
		assert.EqualValues(t, userDomainReturn.Email, userDomain.Email)
		assert.EqualValues(t, userDomainReturn.Password, userDomain.Password)
		assert.EqualValues(t, userDomainReturn.Name, userDomain.Name)
		assert.EqualValues(t, userDomainReturn.Age, userDomain.Age)
	})

	t.Run("when_does_not_exists_an_user_returns_error", func(t *testing.T) {
		email := "test@error.com"
		repository.EXPECT().FindUserByEmail(email).Return(nil, rest_errors.NewNotFoundError("user not found"))

		userDomainReturn, err := service.FindUserByEmailServices(email)

		assert.Nil(t, userDomainReturn)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "user not found")

	})
}

func TestUserDomainService_FindUserByEmailAndPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserPort(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_exists_an_user_returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		email := "test@success.com"
		password := strconv.FormatInt(rand.Int63(), 10)

		userDomain := &domain.UserDomain{
			Email:    email,
			Password: password,
			Name:     "test",
			Age:      90,
			Id:       id,
		}
		repository.EXPECT().FindUserByEmailAndPassword(email, password).Return(userDomain, nil)

		userDomainReturn, err := service.FindUserByEmailAndPasswordServices(email, password)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomainReturn.Id, id)
		assert.EqualValues(t, userDomain.Email, userDomain.Email)
		assert.EqualValues(t, userDomain.Password, userDomain.Password)
		assert.EqualValues(t, userDomain.Name, userDomain.Name)
		assert.EqualValues(t, userDomain.Age, userDomain.Age)
	})
}
