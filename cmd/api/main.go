package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/saad7890/flowship/internal/config"
	"github.com/saad7890/flowship/internal/handler"
	"github.com/saad7890/flowship/internal/infrastructure/postgres"
	"github.com/saad7890/flowship/internal/usecase"
)

func main(){

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	cfg := config.Load()

	db := postgres.New(cfg)

	defer db.Close()

	parcelRepo := postgres.NewParcelRepository(db)
	parcelUsecase := usecase.NewParcelUsecase(parcelRepo)

	router := handler.NewRouter(parcelUsecase)
	fmt.Println("Starting", cfg.AppName)
	fmt.Println("Listening on port:", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}