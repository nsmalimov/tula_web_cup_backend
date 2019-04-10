package sql

import (
	"log"

	"stash.prostream.ru/dsp/sdsp.git/config"
	"stash.prostream.ru/dsp/sdsp.git/helper"
	"stash.prostream.ru/dsp/sdsp.git/interfaces"
	"stash.prostream.ru/dsp/sdsp.git/logging"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func CreateNeededTable(db *sqlx.DB, logger interfaces.LoggerInterface) {
	var sqlCode string

	sqlCode = `
	drop table if exists banners;
	drop table if exists websites;
	drop table if exists banner_templates_fields;
	drop table if exists banner_templates;
	drop table if exists banner_types;
	drop table if exists campaigns;`

	_, err := db.Exec(sqlCode)

	if err != nil {
		logger.Error("Error when try db.Exec (drop), err: %s", err.Error())
	}

	sqlCode = `
	create table websites
	(
    	id   serial primary key,
    	name varchar unique not null,
    	created_on timestamp default now(),
    	updated_on timestamp default now()
	);`
	_, err = db.Exec(sqlCode)

	if err != nil {
		logger.Error("Error when try db.Exec (websites), err: %s", err.Error())
	}

	sqlCode = `
	create table banner_types
	(
    	id         varchar primary key,
		created_on timestamp default now(),
    	updated_on timestamp default now()
	);`
	_, err = db.Exec(sqlCode)

	if err != nil {
		logger.Error("Error when try db.Exec (banner_types), err: %s", err.Error())
	}

	sqlCode = `
	create table banner_templates
	(
    	id               serial primary key,
    	name       		 varchar unique not null,
    	template_content text not null,
    	banner_type_id   varchar unique references banner_types (id),
    	created_on       timestamp default now(),
    	updated_on       timestamp default now(),
    	enable           boolean default true
	);`
	_, err = db.Exec(sqlCode)

	if err != nil {
		logger.Error("Error when try db.Exec (banner_templates), err: %s", err.Error())
	}

	sqlCode = `
	create table banners
	(
    	id             serial primary key,
    	place_id       varchar unique not null,
    	banner_type_id varchar references banner_types (id),
    	website_id     int references websites (id),
    	created_on     timestamp default now(),
    	updated_on     timestamp default now()
	);`
	_, err = db.Exec(sqlCode)

	if err != nil {
		logger.Error("Error when try db.Exec (banners), err: %s", err.Error())
	}

	sqlCode = `
	create table campaigns
	(
    	id         serial primary key,
    	name       varchar unique not null,
    	enable     boolean,
    	created_on timestamp default now(),
    	updated_on timestamp default now()
	);`
	_, err = db.Exec(sqlCode)

	if err != nil {
		logger.Error("Error when try db.Exec (campaigns), err: %s", err.Error())
	}

	sqlCode = `
	-- create table users
	-- (
	--     id serial primary key,
	--     name char
	-- );`
	_, err = db.Exec(sqlCode)

	if err != nil {
		logger.Error("Error when try db.Exec, err: %s", err.Error())
	}

	sqlCode = `
	create table banner_templates_fields
	(
    	id                 serial primary key,
    	banner_template_id int references banner_templates (id),
    	campaign_id        int references campaigns (id),
    	field_name         varchar not null,
    	value              text not null,
    	created_on         timestamp default now(),
    	updated_on         timestamp default now()
	);`
	_, err = db.Exec(sqlCode)

	if err != nil {
		logger.Error("Error when try db.Exec (template_fields), err: %s", err.Error())
	}
}

func main() {
	goLoggingLogger, err := logging.InitAndGetGoLoggingLogger()

	if err != nil {
		log.Printf("Error when try init logger, err: %s", err.Error())
		return
	}

	cfg, err := config.LoadConfig()

	if err != nil {
		goLoggingLogger.Errorf("Error when try init config, err: %s", err.Error())
		return
	}

	if cfg == nil {
		goLoggingLogger.Error("Error when try init config, cfg == nil")
		return
	}

	db, err := helper.InitDB(cfg, goLoggingLogger)

	if err != nil {
		goLoggingLogger.Errorf("Error when try init psql db, %s", err.Error())
		return
	}

	defer func() {
		err := db.Close()

		if err != nil {
			goLoggingLogger.Error("Error when try db.Close, err: %s", err.Error())
		}
	}()

	CreateNeededTable(db, goLoggingLogger)
}
