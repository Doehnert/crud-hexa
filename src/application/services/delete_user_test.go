package services

import (
	"testing"

	"github.com/Doehnert/crud-hexa/src/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserDomainService_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserPort(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_sending_a_valid_userId_returns_success", func(t *testing.T) {
		// Arrange
		id := primitive.NewObjectID().Hex()
		repository.EXPECT().DeleteUser(id).Return(nil)

		// Act
		err := service.DeleteUserService(id)

		// Assert
		assert.Nil(t, err)
	})
}
