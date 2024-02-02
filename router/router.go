

package router

import (
	"toDoList/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// load static files
	r.Static("/static", "./static")
	// load file
	r.LoadHTMLGlob("./templates/*")

	r.GET("/", controller.IndexHandler)

	// v1 group
	v1Group := r.Group("v1")
	{

		v1Group.POST("/todo", controller.CreateHandler)

		v1Group.GET("/todo", controller.GetToDoHandler)

		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})
		v1Group.PUT("/todo/:id", controller.UpdateHandler)
		v1Group.DELETE("/todo/:id", controller.DeleteHandler)
	}

	return r
}
