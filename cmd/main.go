package main

import (
	"github.com/aspandyar/todo-list"
	"github.com/aspandyar/todo-list/pkg/handler"
	"log"
)

func main() {
	srv := new(todo.Server)

	handlers := new(handler.Handler)

	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
