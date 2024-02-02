

package controller

import (
	"net/http"
	"toDoList/model"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*
url -> controller -> logic -> model
*/

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func GetToDoHandler(c *gin.Context) {
	var toDoList []model.Todo
	if err := model.GetTodoList(&toDoList); err != nil {
		// error occur
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		// success
		c.JSON(http.StatusOK, toDoList)
	}
}

func CreateHandler(c *gin.Context) {
	var todo model.Todo
	c.BindJSON(&todo)
	if err := model.CreateTodo(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		// success
		c.JSON(http.StatusOK, todo)
	}
}

func UpdateHandler(c *gin.Context) {
	id, ok := c.Params.Get("id")
	// judge whether have id in request
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "invalid id",
		})
		return
	} else {
		var todo model.Todo
		if err := model.SearchTodo1(id, &todo); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
			return
		} else {
			c.BindJSON(&todo) 
			if err := model.UpdateTodo(&todo); err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			}
		}
	}
}

func DeleteHandler(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "invalid id",
		})
		return
	} else {
		// search id in db
		if err := model.SearchTodo2(id); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
			return
		} else {
			// do delete operation
			if err := model.DeleteTodo(id); err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			}
		}
	}
}
