package repositories

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type MainDbRepository struct {
	DB *sqlx.DB
}

type DbImageData struct {
	Id        int64    `json:"id" db:"id"`
	ImageUrl  string   `json:"token" db:"token"`
	ImageName string   `json:"token" db:"token"`
	UserId    string   `json:"token" db:"token"`
	Rate      float64  `json:"token" db:"token"`
	Tags      []string `json:"token" db:"token"`
}

func (b *MainDbRepository) GetImagesByUserToken(userToken string) (result []DbBannerTemplate, err error) {
	// todo: join
	err = b.DB.Select(&result, `SELECT * FROM banner_templates`)

	if err != nil {
		return nil, err
	}

	return
}
