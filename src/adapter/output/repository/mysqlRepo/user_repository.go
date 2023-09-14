package mysqlrepo

import (
	"database/sql"

	"github.com/Doehnert/crud-hexa/src/application/port/output"
)

const (
	MONGODB_USER_DB = "MONGODB_USER_DB"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(
	db *sql.DB,
) output.UserPort {
	return &userRepository{
		db: db,
	}
}
