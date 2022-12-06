package main

import (
	"employee_manage/config"
	"employee_manage/models"
	"employee_manage/services"
	"encoding/json"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	config.LoadEnv()
	config.DB, _ = config.GormOpen()
	conn, err := amqp.Dial(config.ConfigApp.RabbitMQ.Url)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Println(err)
	}
	defer ch.Close()

	log.Println("Successfully connected to RabbitMQ")
	log.Println("Waiting for messages")

	forever := make(chan bool)

	ConsumerEmail(ch)
	ConsumerResetPassword(ch)
	<-forever
}

func ConsumerResetPassword(ch *amqp.Channel) {
	messages, err := ch.Consume(
		"reset-password", // queue name
		"",               // consumer
		true,             // auto-ack
		false,            // exclusive
		false,            // no local
		false,            // no wait
		nil,              // arguments
	)
	if err != nil {
		log.Println(err)
	}

	go func() {
		for message := range messages {
			go func(message amqp.Delivery) {
				var email string
				json.Unmarshal(message.Body, &email)

				log.Printf(" > Reset password : %s\n", email)

				models.ResetPassword(email, "hoangdz1")
				services.SendMail(email, "Reset password", fmt.Sprintf("New password: %s", "hoangdz1"))
			}(message)
		}
	}()
}

func ConsumerEmail(ch *amqp.Channel) {
	messages, err := ch.Consume(
		"reset-email", // queue name
		"",            // consumer
		true,          // auto-ack
		false,         // exclusive
		false,         // no local
		false,         // no wait
		nil,           // arguments
	)
	if err != nil {
		log.Println(err)
	}

	go func() {
		for message := range messages {
			log.Printf(" > Received message: %s\n", message.Body)

			emails := []string{}
			json.Unmarshal(message.Body, &emails)
			for _, d := range emails {
				config.Publish("reset-password", d)
			}
		}
	}()
}
