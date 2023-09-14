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
	domain := domain.UserDomain{
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
	assert.EqualValues(t, userDomain.Email, domain.Email)
	assert.EqualValues(t, userDomain.Name, domain.Name)
	assert.EqualValues(t, userDomain.Age, domain.Age)
	assert.EqualValues(t, userDomain.Password, domain.Password)
}
