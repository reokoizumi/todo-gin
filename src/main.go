package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func main() {
	



	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))

	r := gin.Default()
	setCors(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(port)
}

func setCors(r *gin.Engine) {
	env := os.Getenv("ENV_NAME")
	var allowOrigins []string
	if env == "prd" {
		allowOrigins = []string{
			"hoge",
		}
	} else {
		allowOrigins = []string{
			os.Getenv("LOCAL_URL"),
		}
	}

	r.Use(cors.New(cors.Config{
		AllowOrigins: allowOrigins,
		AllowMethods: []string{
			"POST",
			"GET",
			"PUT",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))
}
