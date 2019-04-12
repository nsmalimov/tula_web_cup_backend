package controllers

import (
	"net/http"

	"tula_web_cup_backend/app/response"
	"tula_web_cup_backend/repositories/db_repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func CreateUsers(db *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		userToken := ctx.Param("user_token")

		repo := db_repository.DbUsersRepository{
			DB: db,
		}

		err := repo.Create(userToken)

		if err != nil {
			response.Error(err.Error(), http.StatusInternalServerError, ctx)
			return
		}

		resp := response.Response{
			Result: "Created",
		}

		ctx.JSON(http.StatusOK, resp)
	})
}
