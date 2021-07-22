package main

import (
	"log"

	todo_app "github.com/KrisInferno/todo-app"
	"github.com/KrisInferno/todo-app/pkg/handler"
	"github.com/KrisInferno/todo-app/pkg/repository"
	"github.com/KrisInferno/todo-app/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo_app.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
