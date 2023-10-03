// rabbitmq.go
package mq

import (
	"fmt"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

// NewRabbitMQ creates a new RabbitMQ connection.
func NewRabbitMQ(url string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(url)

	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	// With the instance and declare Queues that we can
	// publish and subscribe to.
	_, err = ch.QueueDeclare(
		"fit_queue", // queue name
		true,        // durable
		false,       // auto delete
		false,       // exclusive
		false,       // no wait
		nil,         // arguments
	)

	if err != nil {
		return nil, err
	}

	return &RabbitMQ{conn, ch}, nil
}

// Close closes the RabbitMQ connection.
func (rmq *RabbitMQ) Close() {
	rmq.ch.Close()
	rmq.conn.Close()
}

// Publish sends a message to the RabbitMQ exchange.
func (rmq *RabbitMQ) Publish(routingKey string, data []byte) error {
	err := rmq.ch.Publish(
		"",         // Exchange
		routingKey, // Routing key (queue name)
		false,      // Mandatory
		false,      // Immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
		},
	)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// package mq

// import (
// 	"os"

// 	"github.com/streadway/amqp"
// 	"github.com/tormoder/fit"
// )

// type MessageQuque interface {
// 	sendFitMessage(fit fit.ActivityFile) error
// }

// type RabbitMQ struct {
// 	ch amqp.Channel
// }

// func (mq *RabbitMQ) sendFitMessage(fit *fit.ActivityFile) error {
// 	return nil
// }

// func NewRabbitMq() (*RabbitMQ, error) {

// 	// Define RabbitMQ server URL.
// 	amqpServerURL := os.Getenv("AMQP_SERVER_URL")

// 	// Create a new RabbitMQ connection.
// 	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer connectRabbitMQ.Close()

// 	// Let's start by opening a channel to our RabbitMQ
// 	// instance over the connection we have already
// 	// established.
// 	ch, err := connectRabbitMQ.Channel()
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer ch.Close()

// 	// With the instance and declare Queues that we can
// 	// publish and subscribe to.
// 	_, err = ch.QueueDeclare(
// 		"some test queue", // queue name
// 		true,              // durable
// 		false,             // auto delete
// 		false,             // exclusive
// 		false,             // no wait
// 		nil,               // arguments
// 	)

// 	if err != nil {
// 		panic(err)
// 	}

// 	// publishing a message
// 	err = ch.Publish(
// 		"FitParserQueue", // exchange
// 		"testing",        // key
// 		false,            // mandatory
// 		false,            // immediate
// 		amqp.Publishing{
// 			ContentType: "text/plain",
// 			Body:        []byte("Test Message"),
// 		},
// 	)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// fmt.Println("Successfully published message")
// 	return &RabbitMQ{
// 		ch: ch,
// 	}, nil

// }
