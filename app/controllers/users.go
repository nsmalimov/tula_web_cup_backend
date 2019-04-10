package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
	"tula_web_cup_backend/app/response"
)

func Users(db *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		resp := response.Response{
			Result: "pong",
		}

		ctx.JSON(http.StatusOK, resp)
	})
}
