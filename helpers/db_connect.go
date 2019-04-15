package helpers

import (
	"fmt"
	"time"

	"tula_web_cup_backend/app/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectToPsqlDb(config *config.Configuration) (*sqlx.DB, error) {
	psqlDbName := config.DatabaseName
	psqlDbPassword := config.DatabasePass
	psqlDbUsername := config.DatabaseUser
	psqlDbHost := config.DatabaseHost
	psqlDbPort := config.DatabasePort

	dataSource := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%d sslmode=disable",
		psqlDbName, psqlDbUsername, psqlDbPassword, psqlDbHost, psqlDbPort)

	psqlDbConnect, err := sqlx.Open("postgres", dataSource)

	if err != nil {
		return nil, err
	}

	err = psqlDbConnect.Ping()

	if err != nil {
		return nil, err
	}

	psqlDbConnect.SetConnMaxLifetime(time.Minute * 5)
	psqlDbConnect.SetMaxIdleConns(5)
	psqlDbConnect.SetMaxOpenConns(7)

	return psqlDbConnect, nil
}
