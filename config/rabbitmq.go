package config

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

var ctx = context.TODO()

func Publish() (err error) {
	conn, err := amqp.Dial(ConfigApp.RabbitMQ.Url)
	if err != nil {
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"reset-password-queue", // name
		false,                  // durable
		false,                  // delete when unused
		false,                  // exclusive
		false,                  // no-wait
		nil,                    // arguments
	)
	if err != nil {
		return
	}

	body := "Golang is awesome - Keep Moving Forward!"
	err = ch.PublishWithContext(
		ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	if err != nil {
		return
	}

	return
}
