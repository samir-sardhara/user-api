package repository

import (
	"database/sql"
	"user-api/db/sqlc"
)

type UserRepository struct {
	Queries *sqlc.Queries
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		Queries: sqlc.New(db),
	}
}
