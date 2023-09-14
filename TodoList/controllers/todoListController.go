package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitorpereiraa/go-microservices-starter/TodoList/dtos"
	"github.com/vitorpereiraa/go-microservices-starter/TodoList/services"
)

type TodoListController struct {
	svc services.ITodoListService	
}

func NewTodoListController(todoListService services.ITodoListService) *TodoListController {
	return &TodoListController{
		svc: todoListService,
	}
}

// FindTodoLists godoc
// @Summary      FindTodoLists 
// @Description  Find all todo lists 
// @Tags         todoList
// @Success      200  {object}  []dtos.TodoListDTO
// @Router       /todoList/ [get]
func (t *TodoListController) FindTodoLists(ctx *gin.Context) {
	todoLists := t.svc.FindTodoLists()
	ctx.JSON(http.StatusOK, todoLists)
}

// FindTodoListsByName 	godoc
// @Summary      		FindTodoListByName 
// @Description  		Find a todo list by Name
// @Tags         		todoList
// @Param        		id   path  string  true  "Todo List Name"
// @Success      		200  {object}  dtos.TodoListDTO
// @Failure      		404  {object}  HTTPError  
// @Router       		/todoList/{name} [get]
func (t *TodoListController) FindTodoListByName(ctx *gin.Context) {
	name := ctx.Param("name")

	todoList, err := t.svc.FindTodoListByName(name)

	if err != nil {
		NewError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, todoList)
}

// CreateTodoList	godoc
// @Summary			Create todoList
// @Description		Create a Todo List
// @Tags			todoList
// @Accept			application/json
// @Produce			application/json
// @Param			todoList body dtos.TodoListDTO true "Create todoList"
// @Success			200 {object} dtos.TodoListDTO
// @Failure      	400 {object} HTTPError  
// @Router			/todoList [post]
func (t *TodoListController) CreateTodoList(ctx *gin.Context) {
	var todoListDto dtos.TodoListDTO

	if err := ctx.BindJSON(&todoListDto); err != nil {
		return
	}

	res, err := t.svc.CreateTodoList(todoListDto)

	if err != nil {
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

// DeleteTodoListByName godoc
// @Summary      		DeleteTodoListByName 
// @Description  		Delete a todo list by Name
// @Tags         		todoList
// @Param        		id   path  string  true  "Todo List Name"
// @Success      		200  {object}  dtos.TodoListDTO
// @Failure      		404  {object}  HTTPError  
// @Router       		/todoList/{name} [delete]
func (t *TodoListController) DeleteTodoListByName(ctx *gin.Context) {
	name := ctx.Param("name")

	deletedTodoList, err := t.svc.DeleteTodoListByName(name)

	if err != nil {
		NewError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, deletedTodoList)
}
