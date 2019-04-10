package db

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

func ConnectToPsqlDb() (*sqlx.DB, error) {
	psqlDbName := "postgres"
	psqlDbPassword := "postgres"
	psqlDbUsername := "postgres"
	psqlDbHost := "localhost"
	psqlDbPort := 5432

	dataSource := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%d sslmode=disable",
		psqlDbName, psqlDbUsername, psqlDbPassword, psqlDbHost, psqlDbPort)

	db, err := sqlx.Open("postgres", dataSource)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(7)

}
