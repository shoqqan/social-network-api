package main

import (
	"fmt"
	"log"
	network "social-network-api"
	"social-network-api/pkg/handler"
	"social-network-api/pkg/repository"
	"social-network-api/pkg/service"
)

const (
	envDev  = "dev"
	envProd = "prod"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	fmt.Println("Starting API")
	srv := new(network.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}

}
