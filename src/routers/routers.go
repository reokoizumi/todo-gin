package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/reokoizumi/todo-gin/controllers"
)


func SetupRouter() *gin.Engine {
	r := gin.Default()
	app := r.Group("/app")
	{
		app.GET("todo", controllers.GetTodos)
		app.POST("todo", controllers.CreateATodo)
	}
	return r
}


