package main

import (
	"fmt"
	"log"
	"os"

	"github.com/reokoizumi/todo-gin/config"
	"github.com/reokoizumi/todo-gin/routers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

func main() {
	_, err := gorm.Open(mysql.Open(config.DbURL(config.BuildDBConfig())), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully connect database..")
	}

	err = godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))

	// r := gin.Default()
	// setCors(r)

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	r := routers.SetupRouter()
	r.Run(port)
}
