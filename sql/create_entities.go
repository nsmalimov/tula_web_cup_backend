package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"tula_web_cup_backend/app/config"
)

func CreateNeededTable(db *sqlx.DB) error {
	var sqlCode string

	sqlCode = `
	drop table if exists tags;
	drop table if exists images;
	drop table if exists users;`

	_, err := db.Exec(sqlCode)

	if err != nil {
		return err
	}

	sqlCode = `
	create table users
	(
    	token varchar primary key
	);`
	_, err = db.Exec(sqlCode)

	if err != nil {
		return err
	}

	sqlCode = `
	create table images
	(
    	id          serial primary key,
		image_url   varchar unique not null,
		image_name  varchar not null,
    	user_token  varchar references users (token),
    	resource_id varchar unique not null,
    	rate        float
	);`
	_, err = db.Exec(sqlCode)

	if err != nil {
		return err
	}

	sqlCode = `
	create table tags
	(
    	id       serial primary key,
		tag_name varchar unique not null,
		image_id int references images (id)
	);`
	_, err = db.Exec(sqlCode)

	if err != nil {
		return err
	}

	return nil
}

func main() {
	configApp, err := config.GetConfig()

	if err != nil {
		log.Println(err)
		return
	}

	dbConnect, err := helpers.ConnectToPsqlDb(configApp)

	if err != nil {
		log.Println(err)
		return
	}

	err = CreateNeededTable(dbConnect)

	if err != nil {
		log.Println(err)
	}
}
