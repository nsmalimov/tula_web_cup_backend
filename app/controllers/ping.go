package controllers

import (
	"net/http"

	"tula_web_cup_backend/app/response"

	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	resp := response.Response{
		Result: "pong",
	}

	ctx.JSON(http.StatusOK, resp)
}
