package db_repository

import (
	"github.com/jmoiron/sqlx"
)

type DbImage struct {
	Id         int64   `json:"id" db:"id"`
	ImageUrl   string  `json:"image_url" db:"image_url"`
	ImageName  string  `json:"image_name" db:"image_name"`
	UserToken  string  `json:"user_token" db:"user_token"`
	Tags       []DbTag `json:"tags"`
	ResourceId string  `json:"resource_id" db:"resource_id"`

	// todo: omitempty разобраться
	Rate float64 `json:"rate" db:"rate"`
}

type DbImagesRepository struct {
	DB *sqlx.DB
}

func (b *DbImagesRepository) GetAll() ([]DbImage, error) {
	var dbImages []DbImage

	err := b.DB.Select(&dbImages, `SELECT * FROM images;`)

	if err != nil {
		return nil, err
	}

	return dbImages, nil
}

func (b *DbImagesRepository) GetImageById(imageId int64) (*DbImage, error) {
	var image DbImage

	err := b.DB.Get(&image, `SELECT * FROM images WHERE id = $1;`, imageId)

	if err != nil {
		return nil, err
	}

	return &image, nil
}

func (b *DbImagesRepository) GetImagesByUserToken(userToken string) ([]DbImage, error) {
	// todo: join

	var dbImages []DbImage

	query := `SELECT * FROM images WHERE user_token = $1`

	err := b.DB.Select(&dbImages, query, userToken)

	if err != nil {
		return nil, err
	}

	return dbImages, nil
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

func (b *DbImagesRepository) InsertMany(dbImages []DbImage) error {
	tx := b.DB.MustBegin()

	for _, dbImage := range dbImages {
		_, err := tx.NamedExec("INSERT INTO images "+
			"(image_url, image_name, user_token, rate, resource_id) VALUES "+
			"(:image_url, :image_name, :user_token, :rate, :resource_id)",
			&dbImage)

		if err != nil {
			return err
		}
	}
	err := tx.Commit()

	if err != nil {
		return err
	}

	return err
}

func (b *DbImagesRepository) DeleteByimageIds(dbImageIds []int64) error {
	tx := b.DB.MustBegin()

	for _, dbImageId := range dbImageIds {
		_ = tx.MustExec(`DELETE FROM images WHERE id=$1`, &dbImageId)
	}

	err := tx.Commit()

	if err != nil {
		return err
	}

	return err
}
