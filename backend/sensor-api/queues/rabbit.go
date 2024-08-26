package queues

import (
	"context"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var queueName string
var q amqp.Queue
var Ch *amqp.Channel

// Initialize MQ publisher
func InitRabbitMQ() func() {
	mqUrl := os.Getenv("RABBITMQ_URL")
	queueName = os.Getenv("QUEUE_NAME")

	conn, err := amqp.Dial(mqUrl)
	failOnError(err, "Failed to connect to Rabbitmq")
	// defer conn.Close()

	Ch, err = conn.Channel()
	failOnError(err, "Failed to open a channel")
	// defer ch.Close()

	q, err = Ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare queue")

	return func() {
		// Cleaning up queue resources
		defer conn.Close()
		defer Ch.Close()
	}
}

func PublishMessage(msg []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return Ch.PublishWithContext(
		ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		},
	)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
