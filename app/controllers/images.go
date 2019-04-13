package controllers

import (
	"fmt"
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
		log.Println("Request: UpdateImages")

		var imagesUpdateRequest ImagesUpdateRequest

		// todo: сделать на джоинах для практики
		userToken := ctx.Param("user_token")

		log.Printf("userToken: %s", userToken)

		err := ctx.BindJSON(&imagesUpdateRequest)

		if err != nil {
			s := fmt.Sprintf("Error when try ctx.BindJSON, err: %s", err)
			response.Error(s, http.StatusBadRequest, ctx)
			return
		}

		log.Printf("request need update images num: %d", len(imagesUpdateRequest.DbImages))

		repoImages := db_repository.DbImagesRepository{
			DB: db,
		}

		resp := response.Response{
			Result: "All actual",
		}

		imagesByUserToken, err := repoImages.GetImagesByUserToken(userToken)

		if err != nil {
			s := fmt.Sprintf("Error when try repoImages.GetImagesByUserToken, err: %s", err)
			response.Error(s, http.StatusInternalServerError, ctx)
			return
		}

		log.Printf("images exist by user: %d", len(imagesByUserToken))

		imagesByUserTokenMap := make(map[string]db_repository.DbImage)

		for _, imageByUserToken := range imagesByUserToken {
			imagesByUserTokenMap[imageByUserToken.ResourceId] = imageByUserToken
		}

		imagesUpdateRequestMap := make(map[string]db_repository.DbImage)

		for _, imageUpdateRequest := range imagesUpdateRequest.DbImages {
			imagesUpdateRequestMap[imageUpdateRequest.ResourceId] = imageUpdateRequest
		}

		// ------ добавление
		var imagesNeedCreate []db_repository.DbImage

		for _, dbImageFromRequest := range imagesUpdateRequest.DbImages {
			if _, ok := imagesByUserTokenMap[dbImageFromRequest.ResourceId]; ok {
				// image есть в базе, проверка на апдейт
			} else {
				imagesNeedCreate = append(imagesNeedCreate, db_repository.DbImage{
					ImageUrl:   dbImageFromRequest.ImageUrl,
					ImageName:  dbImageFromRequest.ImageName,
					UserToken:  userToken,
					ResourceId: dbImageFromRequest.ResourceId,

					// todo: разобрться как сделать без 0
					Rate: -1,
				})
			}
		}

		log.Printf("imagesNeedCreate num: %d", len(imagesNeedCreate))

		repoUsers := db_repository.DbUsersRepository{
			DB: db,
		}

		user := repoUsers.GetUserByToken(userToken)

		if user.Token == "" {
			fmt.Printf("User not exist, will create, %s\n", userToken)
			err = repoUsers.Create(userToken)

			if err != nil {
				s := fmt.Sprintf("Error when try repoUsers.Create, err: %s", err)
				response.Error(s, http.StatusInternalServerError, ctx)
				return
			}
		}

		if len(imagesNeedCreate) != 0 {
			err = repoImages.InsertMany(imagesNeedCreate)

			if err != nil {
				s := fmt.Sprintf("Error when try repoImages.InsertMany, err: %s", err)
				response.Error(s, http.StatusInternalServerError, ctx)
				return
			}

			resp.Result = "Updated"
		}

		// ------ удаление
		var imageIdsNeedDelete []int64

		for _, imageByUserToken := range imagesByUserToken {
			if _, ok := imagesUpdateRequestMap[imageByUserToken.ResourceId]; ok {
				// pass
			} else {
				imageIdsNeedDelete = append(imageIdsNeedDelete, imageByUserToken.Id)
			}
		}

		log.Printf("imageIdsNeedDelete num: %d", len(imageIdsNeedDelete))

		err = repoImages.DeleteByimageIds(imageIdsNeedDelete)

		if err != nil {
			s := fmt.Sprintf("Error when try repoImages.DeleteByimageIds, err: %s", err)
			response.Error(s, http.StatusInternalServerError, ctx)
			return
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
		log.Println("Request: GetAllImages")

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
