package repositories

import (
	"github.com/jmoiron/sqlx"
)

type MainDbRepository struct {
	DB *sqlx.DB
}
