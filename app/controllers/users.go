package controllers

import (
	"log"
	"net/http"

	"tula_web_cup_backend/app/response"
	"tula_web_cup_backend/repositories/db_repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func CreateUsers(db *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		log.Println("Request: CreateUsers")

		userToken := ctx.Param("user_token")

		repo := db_repository.DbUsersRepository{
			DB: db,
		}

		resp := response.Response{}

		user := repo.GetUserByToken(userToken)

		if user.Token != "" {
			resp.Result = "User in db already exist"
		} else {
			err := repo.Create(userToken)

			if err != nil {
				response.Error(err.Error(), http.StatusInternalServerError, ctx)
				return
			}

			resp.Result = "User in db created"
		}

		ctx.JSON(http.StatusOK, resp)
	})
}
