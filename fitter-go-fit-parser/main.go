package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	storage "github.com/notaduck/fitter-go-fit-parser/db"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/tormoder/fit"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {

	s, err := storage.NewPostgresStore()

	if err != nil {
		fmt.Println(err)
	}

	conn, err := amqp.Dial("amqp://consumer:consumer@rabbitmq:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// opening a channel over the connection established to interact with RabbitMQ
	ch, err := conn.Channel()

	if err != nil {
		panic(err)
	}

	defer ch.Close()

	fmt.Println("Successfully connected to RabbitMQ instance")

	msgs, err := ch.Consume(
		"fit_queue", // queue
		"",          // consumer
		// true,          // auto ack
		true,  // auto ack
		false, // exclusive
		false, // no local
		false, // no wait
		nil,   //args
	)
	if err != nil {
		panic(err)
	}

	forever := make(chan bool)
	go func() {
		for msg := range msgs {

			fmt.Println("Message recieved")

			type FitterBody struct {
				FitFile []byte
				UserId  int
			}

			body := FitterBody{}

			json.Unmarshal(msg.Body, &body)

			reader := bytes.NewReader(body.FitFile)

			fmt.Println("decode fit")
			fit, err := fit.Decode(reader)
			if err != nil {
				fmt.Println("error 1")
			}

			activity, err := fit.Activity()

			if err != nil {
				fmt.Println("error 2")
			}
			aId, err := s.CreateActivity(body.UserId, activity)

			fmt.Println("ID", aId)

			if err != nil {
				fmt.Println("error 3")
			}

			fmt.Println("insert activities")
			err = s.CreateMsgRecords(aId, activity.Records)

			if err != nil {
				fmt.Println(err)
				panic(err)
			}

		}
	}()

	fmt.Println("Waiting for messages...")
	<-forever
}
