package routers

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/reokoizumi/todo-gin/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	setCors(r)
	app := r.Group("/app")
	{
		app.GET("todo", controllers.GetTodos)
		app.POST("todo", controllers.CreateATodo)
	}
	return r
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
