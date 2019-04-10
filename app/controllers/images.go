package controllers

import (
	"net/http"
	"strconv"
	"tula_web_cup_backend/repositories/db_repository"

	"tula_web_cup_backend/app/response"
	"tula_web_cup_backend/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type ImagesUpdateRequest struct {
	ImageUrls []string `json:"images_url"`
}

func ImagesUpdate(db *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var imagesUpdateRequest ImagesUpdateRequest

		userToken, err := strconv.Atoi(ctx.Param("user_token"))

		if err != nil {
			response.Error(err.Error(), http.StatusBadRequest, ctx)
		}

		err = ctx.BindJSON(&imagesUpdateRequest)

		if err != nil {
			response.Error(err.Error(), http.StatusBadRequest, ctx)
		}

		repo := repositories.MainDbRepository{
			DB: db,
		}

		err = repo.GetImagesByUserToken(dbUser)

		// проверка, удаление и добавление

		if err != nil {
			response.Error(err, http.StatusInternalServerError, ctx)
		}

		resp := response.Response{
			Result: "Created",
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
