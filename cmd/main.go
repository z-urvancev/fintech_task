package main

import (
	"fintech/config"
	"fintech/internal/handler"
	repository "fintech/internal/repository/impl"
	"fintech/internal/router"
	usecase "fintech/internal/usecase/impl"
	"flag"
	"log"
)

func main() {
	var cfg config.Config
	configErr := config.InitConfig(&cfg)
	if configErr != nil {
		log.Fatalln(configErr)
	}

	var storeType string
	flag.StringVar(&storeType, "store_type", "inMemory", "store type")
	flag.Parse()
	repo, repoErr := repository.InitRepository(storeType, &cfg)
	if repoErr != nil {
		log.Fatalln(repoErr)
	}

	useCase := usecase.NewUseCase(repo)

	handler := handler.NewHandler(useCase)

	router := router.InitRoutes(handler)

	runErr := router.Run(cfg.Port)
	if runErr != nil {
		log.Fatal(runErr)
	}
}
