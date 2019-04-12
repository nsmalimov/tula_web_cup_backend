package controllers

import (
	"log"
	"net/http"

	"tula_web_cup_backend/app/response"
	"tula_web_cup_backend/repositories/db_repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func CreateTag(db *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var dbTag db_repository.DbTag

		log.Println("Request to create tag")

		err := ctx.BindJSON(&dbTag)

		if err != nil {
			response.Error(err.Error(), http.StatusBadRequest, ctx)
		}

		repo := db_repository.DbTagsRepository{
			DB: db,
		}

		err = repo.Create(dbTag)

		if err != nil {
			response.Error(err.Error(), http.StatusInternalServerError, ctx)
		}

		resp := response.Response{
			Result: "Created",
		}

		ctx.JSON(http.StatusOK, resp)
	})
}
