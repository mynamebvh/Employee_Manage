package config

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

var ctx = context.TODO()

func Publish(queue string, body interface{}) (err error) {
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
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		fmt.Println("Error queue declare: ", err)
		return
	}

	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err = encoder.Encode(body)

	if err != nil {
		fmt.Println("Error encoder json: ", err)
		return
	}

	err = ch.PublishWithContext(
		ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        b.Bytes(),
		})

	if err != nil {
		return
	}

	return
}
