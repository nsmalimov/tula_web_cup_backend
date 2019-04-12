package controllers

import (
	"log"
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

		log.Printf("userToken: %s", userToken)

		err := ctx.BindJSON(&imagesUpdateRequest)

		log.Printf("request need update images num: %d", len(imagesUpdateRequest.DbImages))

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

		log.Printf("images exist by user: %d", len(imagesByUserToken))

		imagesByUserTokenMap := make(map[string]db_repository.DbImage)

		for _, imageByUserToken := range imagesByUserToken {
			imagesByUserTokenMap[imageByUserToken.ImageUrl] = imageByUserToken
		}

		if err != nil {
			response.Error(err.Error(), http.StatusInternalServerError, ctx)
			return
		}

		// проверка, удаление и добавление

		var imagesNeedCreate []db_repository.DbImage

		for _, dbImageFromRequest := range imagesUpdateRequest.DbImages {
			if _, ok := imagesByUserTokenMap[dbImageFromRequest.ImageUrl]; ok {
				// image есть в базе, проверка на апдейт
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

		log.Printf("imagesNeedCreate num: %d", len(imagesNeedCreate))

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
		log.Println("Request to get all images")

		repo := db_repository.DbImagesRepository{
			DB: db,
		}

		images, err := repo.GetAll()

		repoTags := db_repository.DbTagsRepository{
			DB: db,
		}

		// todo: slowly
		for index, image := range images {
			tags, err := repoTags.GetTagsByImageId(image.Id)

			if err != nil {
				log.Println(err)
				continue
			}

			images[index].Tags = tags
		}

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
		tagName := ctx.Param("tag_name")

		log.Printf("tagName: %s", tagName)

		repoTags := db_repository.DbTagsRepository{
			DB: db,
		}

		imagesIds, err := repoTags.GetImageIdsByTagName(tagName)

		if err != nil {
			response.Error(err.Error(), http.StatusInternalServerError, ctx)
		}

		repoImages := db_repository.DbImagesRepository{
			DB: db,
		}

		var images []*db_repository.DbImage

		for _, imageId := range imagesIds {
			image, err := repoImages.GetImageById(imageId)

			tags, err := repoTags.GetTagsByImageId(imageId)

			if err != nil {
				log.Println(err)
				continue
			}

			image.Tags = tags

			images = append(images, image)
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
