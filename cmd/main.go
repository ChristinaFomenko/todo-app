package main

import (
	"github.com/KrisInferno/todo-app"
	"log"
)

func main() {
	srv := new(todo_app.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
