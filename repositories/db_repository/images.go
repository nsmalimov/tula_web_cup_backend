package db_repository

import "github.com/jmoiron/sqlx"

type DbImage struct {
	Id        int64   `json:"id" db:"id"`
	ImageUrl  string  `json:"images_url" db:"image_url"`
	ImageName string  `json:"image_name" db:"image_name"`
	UserToken string  `json:"user_token" db:"user_token"`
	Rate      float64 `json:"rate,omitempty" db:"rate"`
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

func (b *DbImagesRepository) GetImagesByUserToken(userToken string) ([]DbImage, error) {
	// todo: join

	var dbImages []DbImage

	query := `SELECT * FROM images WHERE user_token = $1`

	err := b.DB.Get(dbImages, query, userToken)

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
			"(image_url, image_name, user_id, rate) VALUES "+
			"(:image_url, :image_name, :user_id, :rate)",
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
