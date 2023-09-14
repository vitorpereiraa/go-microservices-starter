package main

import (
	"log"
	"os"

	"github.com/vitorpereiraa/go-microservices-starter/TodoList/config"
	"github.com/vitorpereiraa/go-microservices-starter/TodoList/controllers"
	_ "github.com/vitorpereiraa/go-microservices-starter/TodoList/docs"
	"github.com/vitorpereiraa/go-microservices-starter/TodoList/repositories"
	"github.com/vitorpereiraa/go-microservices-starter/TodoList/router"
	"github.com/vitorpereiraa/go-microservices-starter/TodoList/services"
)

func main() {
	settings, err := config.LoadSettings()

	if err != nil {
		log.Fatalf("Fatal error loading the settings file: %v", err)
		os.Exit(1)
	}

	db, err := config.LoadDB(settings)
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
		os.Exit(1)
	}

	repo := repositories.NewTodoListRepositoryPostgres(db)
	service := services.NewTodoListService(repo)
	controller := controllers.NewTodoListController(service) 
	router := router.NewRouter(controller)

	err = router.Run(settings.SERVER_HOST + ":" + settings.SERVER_PORT)

	if err != nil {
		log.Fatal("Server failed to start")
		os.Exit(1)
	}
}
