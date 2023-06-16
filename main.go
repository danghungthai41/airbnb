package main

import (
	"airbnb-golang/config"
	placehttp "airbnb-golang/internal/places/delivery/http"
	placerepository "airbnb-golang/internal/places/repository"
	placeusecase "airbnb-golang/internal/places/usecase"
	"airbnb-golang/pkg/db/mysql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	filename := "config/config-local.yml"
	if os.Getenv("env") == "production" {
		filename = "config/config-prod.yml"
	}
	cfg, err := config.LoadConfig(filename)
	if err != nil {
		log.Fatalln("LoadConfig:", err)
	}
	db, err := mysql.NewMySQL(cfg)
	if err != nil {
		log.Fatal("Cannot connect to mysql", err)
	}
	fmt.Println(db)

	r := gin.Default()

	placerepo := placerepository.NewPlaceRepo(db)
	placeuc := placeusecase.NewPlaceUC(placerepo)
	placehdl := placehttp.NewPlaceHandler(placeuc)
	v1 := r.Group("api/v1")
	{
		v1.POST("/places", placehdl.CreatePlace())
		v1.GET("/places", placehdl.GetPlaces())
		v1.GET("/places/:id", placehdl.GetPlaceByID())
		v1.DELETE("/places/:id", placehdl.DeletePlace())
		v1.PUT("/places/:id", placehdl.UpdatePlace())

	}

	r.Run(":" + cfg.App.Port)

}
