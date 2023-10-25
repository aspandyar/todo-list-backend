package main

import (
	"github.com/aspandyar/todo-list"
	"github.com/aspandyar/todo-list/pkg/handler"
	"github.com/aspandyar/todo-list/pkg/repository"
	"github.com/aspandyar/todo-list/pkg/service"
	"log"
)

func main() {
	srv := new(todo.Server)

	repo := repository.NewRepository()
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
