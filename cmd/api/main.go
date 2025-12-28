package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/saad7890/flowship/internal/config"
)

func main(){

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	cfg := config.Load()
	fmt.Println("Starting", cfg.AppName)
	fmt.Println("Listening on port:", cfg.Port)
	
}