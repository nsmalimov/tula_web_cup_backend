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

func (b *DbTagsRepository) GetTagsByImageId(imageId int64) ([]DbTag, error) {
	var dbTags []DbTag

	query := `SELECT * FROM tags WHERE image_id = $1`

	err := b.DB.Select(&dbTags, query, imageId)

	if err != nil {
		return nil, err
	}

	return dbTags, nil
}

func (b *DbTagsRepository) GetImageIdsByTagName(tagName string) ([]int64, error) {
	var imageIds []int64

	query := `SELECT image_id FROM tags WHERE tag_name = $1`

	err := b.DB.Select(&imageIds, query, tagName)

	if err != nil {
		return nil, err
	}

	return imageIds, nil
}
