package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	storage "github.com/notaduck/fitter-go/db"
	"github.com/notaduck/fitter-go/handler"
	"github.com/notaduck/fitter-go/mq"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}

	storage, err := storage.NewPostgresStore()

	if err != nil {
		log.Fatal(err)
	}

	if err := storage.Init(); err != nil {
		log.Fatal(err)
	}

	mqUrl := os.Getenv("AMQP_SERVER_URL")

	// Create a new RabbitMQ connection
	rmq, err := mq.NewRabbitMQ(mqUrl)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer rmq.Close()

	server := handler.NewAPIServer(":3030", storage, *rmq)

	server.Run()
}
