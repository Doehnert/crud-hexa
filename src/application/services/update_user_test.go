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

func TestUserDomainService_UpdateUserService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserPort(ctrl)
	services := NewUserDomainService(repository)

	t.Run("when_sendin_a_valid_user_dn_userId_returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		userDomain := &domain.UserDomain{
			Name: "test",
			Age:  90,
		}
		repository.EXPECT().UpdateUser(id, *userDomain).Return(nil)

		// Act
		err := services.UpdateUserService(id, *userDomain)

		// Assert
		assert.Nil(t, err)
	})

	t.Run("when_update_user_repo_fails_return_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		userDomain := &domain.UserDomain{
			Name: "test",
			Age:  49,
		}
		repository.EXPECT().UpdateUser(id, *userDomain).Return(rest_errors.NewInternalServerError("error trying to update user"))

		// Act
		err := services.UpdateUserService(id, *userDomain)

		// Assert
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "error trying to update user")
	})
}
