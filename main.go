package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"tula_web_cup_backend/app/controllers"
	"tula_web_cup_backend/db"

	"github.com/gin-gonic/gin"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	router := gin.New()

	psqlDbConnect, err := db.ConnectToPsqlDb()

	router.GET("/ping", controllers.Ping)

	router.POST("/user", controllers.Users(db))

	portStart := 9090

	err := router.Run(fmt.Sprintf(":%d", portStart))

	if err != nil {
		log.Print(err)
	}
}
