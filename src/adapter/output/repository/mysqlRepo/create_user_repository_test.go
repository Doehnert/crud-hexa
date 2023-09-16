package mysqlrepo

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Doehnert/crud-hexa/src/application/domain"
	"github.com/stretchr/testify/assert"
)

func TestUserRepoMysql_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database connection")
	}
	defer db.Close()

	userRepository := NewUserRepository(db)
	domain := &domain.UserDomain{
		Id:       "id",
		Email:    "test@test.com",
		Name:     "test",
		Password: "test",
		Age:      90,
	}

	// Set up expectations for database queries using the mock object
	mock.ExpectExec("INSERT INTO users").WithArgs().WillReturnResult(sqlmock.NewResult(1, 1))

	// Act
	userDomain, err := userRepository.CreateUser(domain)

	// Asserts
	assert.Nil(t, err)
	assert.EqualValues(t, userDomain.GetEmail(), domain.GetEmail())
	assert.EqualValues(t, userDomain.GetName(), domain.GetName())
	assert.EqualValues(t, userDomain.GetAge(), domain.GetAge())
	assert.EqualValues(t, userDomain.GetPassword(), domain.GetPassword())
}
