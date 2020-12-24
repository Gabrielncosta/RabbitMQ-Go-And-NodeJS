package main

import (
	"github.com/streadway/amqp"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
			log.Printf("No .env file found")
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf(msg);
	}
}

func main() {
	connectionString, ok := os.LookupEnv("CONNECTION_STRING")

	if !ok {
		log.Printf("%s not set\n", connectionString)
	}

	conn, err := amqp.Dial(connectionString)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to create channel")
	defer ch.Close()

	msgs, err := ch.Consume(
		"test-queue",
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to consume queue")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("MESSAGE: %s", d.Body)
			d.Ack(true)
		}
	}()

	log.Printf("[*] Waiting for logs. To Exit press CTRL + C")
	<-forever
}