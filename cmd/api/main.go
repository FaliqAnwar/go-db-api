package main

import (
	"go-db-api/internal/repository"
	"go-db-api/internal/utils/confloader"
	"log"
)

func main() {
	cfg, err := confloader.MustLoad("go-db-api")
	if err != nil {
		panic(err)
	}
	log.Println(cfg)

	repo := repository.NewRepository(cfg.DB)
	log.Println(repo)
}
