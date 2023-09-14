package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/vitorpereiraa/go-microservices-starter/TodoList/controllers"
)

// @title Todo List API
// @version v1
// @description This is a template REST API written in golang using Onion Architecture

// @contact.name   Vitor Pereira
// @contact.email  vitor.cbs.pereira@gmail.com

// @license.name  WTFPL license
// @license.url   http://www.wtfpl.net/

// @host      localhost:8080
// @BasePath  /api
func NewRouter(ctrl controllers.ITodoListController) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")

	todoListRouter := api.Group("/todoList") 
	{
		todoListRouter.GET("/", ctrl.FindTodoLists)
		todoListRouter.GET("/:name", ctrl.FindTodoListByName)
		todoListRouter.POST("/", ctrl.CreateTodoList)
		todoListRouter.DELETE("/:name", ctrl.DeleteTodoListByName)
	}

	return r
}
