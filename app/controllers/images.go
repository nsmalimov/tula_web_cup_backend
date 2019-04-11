package controllers

import (
	"fmt"
	"net/http"
	"tula_web_cup_backend/app/response"
	"tula_web_cup_backend/repositories/db_repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type ImagesUpdateRequest struct {
	DbImages []db_repository.DbImage `json:"images"`
}

func UpdateImages(db *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var imagesUpdateRequest ImagesUpdateRequest

		// todo: сделать на джоинах для практики
		userToken := ctx.Param("user_token")

		fmt.Println(userToken)

		err := ctx.BindJSON(&imagesUpdateRequest)

		fmt.Println(imagesUpdateRequest)

		if err != nil {
			response.Error(err.Error(), http.StatusBadRequest, ctx)
			return
		}

		repo := db_repository.DbImagesRepository{
			DB: db,
		}

		resp := response.Response{
			Result: "All actual",
		}

		imagesByUserToken, err := repo.GetImagesByUserToken(userToken)

		imagesByUserTokenMap := make(map[string]db_repository.DbImage)

		for _, dbImageFromDb := range imagesByUserToken {
			imagesByUserTokenMap[dbImageFromDb.ImageUrl] = dbImageFromDb
		}

		if err != nil {
			response.Error(err.Error(), http.StatusInternalServerError, ctx)
			return
		}

		// проверка, удаление и добавление

		var imagesNeedCreate []db_repository.DbImage

		for _, dbImageFromRequest := range imagesUpdateRequest.DbImages {
			if _, ok := imagesByUserTokenMap[dbImageFromRequest.ImageUrl]; ok {

			} else {
				imagesNeedCreate = append(imagesNeedCreate, db_repository.DbImage{
					ImageUrl:  dbImageFromRequest.ImageUrl,
					ImageName: dbImageFromRequest.ImageName,
					UserToken: userToken,

					// todo: разобрться как сделать без 0
					Rate: -1,
				})
			}

		}

		if len(imagesNeedCreate) != 0 {
			err = repo.InsertMany(imagesNeedCreate)

			if err != nil {
				response.Error(err.Error(), http.StatusInternalServerError, ctx)
				return
			}

			resp.Result = "Updated"
		}

		ctx.JSON(http.StatusOK, resp)
	})
}

func RateImage(db *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {

	})
}

func GetAllImages(db *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		repo := db_repository.DbImagesRepository{
			DB: db,
		}

		images, err := repo.GetAll()

		if err != nil {
			response.Error(err.Error(), http.StatusInternalServerError, ctx)
		}

		resp := response.Response{
			Result: images,
		}

		ctx.JSON(http.StatusOK, resp)
	})
}

func GetAllImagesByTag(db *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		repo := db_repository.DbImagesRepository{
			DB: db,
		}

		images, err := repo.GetAll()

		if err != nil {
			response.Error(err.Error(), http.StatusInternalServerError, ctx)
		}

		resp := response.Response{
			Result: images,
		}

		ctx.JSON(http.StatusOK, resp)
	})
}

func GetAllSortedImages(db *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		sortParam := ctx.Param("sort_param")

		if sortParam == "" || (sortParam != "name" && sortParam != "rate") {
			response.Error("Sort param empty or not one of (name, rate)", http.StatusBadRequest, ctx)
		}

		repo := db_repository.DbImagesRepository{
			DB: db,
		}

		images, err := repo.GetAllSortedImages(sortParam)

		if err != nil {
			response.Error(err.Error(), http.StatusInternalServerError, ctx)
		}

		resp := response.Response{
			Result: images,
		}

		ctx.JSON(http.StatusOK, resp)
	})
}
