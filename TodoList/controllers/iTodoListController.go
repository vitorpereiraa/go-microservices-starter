package controllers

import (
	"github.com/gin-gonic/gin"
)

type ITodoListController interface {
	FindTodoLists(ctx *gin.Context)
	FindTodoListByName(ctx *gin.Context)
	CreateTodoList(ctx *gin.Context)
	DeleteTodoListByName(ctx *gin.Context)
}
