package db_repository

import (
	"github.com/jmoiron/sqlx"
)

type DbTag struct {
	Id      int64  `json:"id" db:"id"`
	TagName string `json:"tag_name" db:"tag_name"`
	ImageId string `json:"image_id" db:"image_id"`
}

type DbTagsRepository struct {
	DB *sqlx.DB
}

func (b *DbTagsRepository) Create(dbTag DbTag) error {
	query := `INSERT INTO tags (tag_name, image_id) VALUES ($1, $2)`

	_, err := b.DB.Query(query, dbTag.TagName, dbTag.ImageId)

	return err
}
