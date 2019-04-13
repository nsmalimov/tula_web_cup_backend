package db_repository

import (
	"github.com/jmoiron/sqlx"
)

type DbUser struct {
	Token string `json:"token" db:"token"`
}

type DbUsersRepository struct {
	DB *sqlx.DB
}

func (b *DbUsersRepository) Create(userToken string) error {
	query := `INSERT INTO users (token) VALUES ($1)`

	_, err := b.DB.Query(query, userToken)

	return err
}

func (b *DbUsersRepository) GetUserByToken(userToken string) *DbUser {
	var dbUser DbUser

	query := `SELECT * FROM users WHERE token = $1`

	// err if only user not exist
	_ = b.DB.Get(&dbUser, query, userToken)

	return &dbUser
}
