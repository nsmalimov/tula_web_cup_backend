package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"tula_web_cup_backend/app/config"
	"tula_web_cup_backend/app/controllers"
	"tula_web_cup_backend/helpers"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	router := gin.New()

	configDir := flag.String("config_dir", "", "config path")

	flag.Parse()

	configApp, err := config.GetConfig(*configDir)

	if err != nil {
		log.Println(err)
		return
	}

	// todo: защитить ручки от неавторизованного удаления (токен доступа?)

	psqlDbConnect, err := helpers.ConnectToPsqlDb(configApp)

	if err != nil {
		log.Println(err)
		return
	}

	httpClient := &http.Client{}

	router.Use(cors.Default())

	router.GET("/ping", controllers.Ping)

	// todo: PUT, DELETE
	// +
	router.POST("/users/:user_token", controllers.CreateUsers(psqlDbConnect))

	// todo: PUT, DELETE
	// +
	router.POST("/tags", controllers.CreateTag(psqlDbConnect))

	// приходит юзер, мы апдейтим базу, забираем все его картинки (много)
	// +
	router.POST("/images/:user_token", controllers.UpdateImages(psqlDbConnect))

	router.GET("/images_update_urls", controllers.UpdateImageUrls(psqlDbConnect, httpClient))

	// оценить картинку image_id=int rate=float
	router.GET("/rate", controllers.RateImage(psqlDbConnect))

	// забрать все картинки (общие)
	// +

	// todo: подумать, возможен ли апдейт
	router.GET("/images", controllers.GetAllImages(psqlDbConnect))

	// todo: забрать все картинки только юзера

	//router.GET("/images", controllers.GetAllImagesByUserToken(psqlDbConnect))

	router.GET("/images_by_tag_name/:tag_name", controllers.GetAllImagesByTag(psqlDbConnect))

	// todo: сортировка по возрастанию

	// name, rate
	// +
	router.GET("/images_sort/:sort_param", controllers.GetAllSortedImages(psqlDbConnect))

	portStart := configApp.PortStart

	err = router.Run(fmt.Sprintf(":%d", portStart))

	if err != nil {
		log.Println(err)
	}
}
