package db_repository

import (
	"github.com/jmoiron/sqlx"
)

type DbUser struct {
	Id    int64  `json:"id" db:"id"`
	Token string `json:"token" db:"token"`
}

type DbUsersRepository struct {
	DB *sqlx.DB
}

func (b *DbUsersRepository) Create(dbUser DbUser) error {
	query := `INSERT INTO users (token) VALUES ($1)`

	_, err := b.DB.Query(query, dbUser.Token)

	return err
}
