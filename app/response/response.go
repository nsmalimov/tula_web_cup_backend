package response

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Result interface{} `json:"result"`
}

func Error(errorText string, code int, ctx *gin.Context) {
	ctx.Status(code)
	log.Println(errorText)

	_, err := ctx.Writer.Write([]byte(errorText))

	if err != nil {
		log.Println(err)
	}
}
