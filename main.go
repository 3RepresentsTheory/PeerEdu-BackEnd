package main

import (
	"PeerEdu-BackEnd/database"
	"PeerEdu-BackEnd/router"
	"PeerEdu-BackEnd/util/config"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	config.Init()
	database.Init()

	if config.Config.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	app := gin.Default()
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Content-Type", "fake-cookie", "webvpn-cookie"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		MaxAge:       12 * time.Hour,
	}))
	router.SetRouter(app)
	fmt.Println("PeerEdu will run on port " + config.Config.Port)
	err := app.Run(":" + config.Config.Port)
	if err != nil {
		panic(err)
	}
}
