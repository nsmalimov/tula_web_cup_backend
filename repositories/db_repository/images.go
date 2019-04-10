package db_repository

import "github.com/jmoiron/sqlx"

type DbImage struct {
	Id        int64  `json:"id" db:"id"`
	ImageUrl  string `json:"token" db:"token"`
	ImageName string `json:"token" db:"token"`
	UserId    string `json:"token" db:"token"`
	Rate      int64  `json:"token" db:"token"`
}

type DbImagesRepository struct {
	DB *sqlx.DB
}

func (b *DbImagesRepository) GetAll() (result []DbImage, err error) {
	err = b.DB.Select(&result, `SELECT * FROM images`)

	if err != nil {
		return nil, err
	}

	return
}

func (b *DbImagesRepository) GetAllSortedImages(sortParam string) (*[]DbImage, error) {
	var dbImages []DbImage

	query := `SELECT * FROM images ORDER BY $1 DESC'`

	err := b.DB.Get(dbImages, query, sortParam)

	if err != nil {
		return nil, err
	}

	return &dbImages, nil
}
