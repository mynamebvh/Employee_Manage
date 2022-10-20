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

	ConsumerExcel(ch)
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
				var data services.DataQueue
				json.Unmarshal(message.Body, &data)

				log.Printf(" > Reset password : %s\n", data)

				models.ResetPassword(data.Email, "hoangdz1")
				services.SendMail(data.Email, "Reset password", fmt.Sprintf("New password: %s", "hoangdz"))
			}(message)

		}
	}()
}

func ConsumerExcel(ch *amqp.Channel) {
	messages, err := ch.Consume(
		"excel", // queue name
		"",      // consumer
		true,    // auto-ack
		false,   // exclusive
		false,   // no local
		false,   // no wait
		nil,     // arguments
	)
	if err != nil {
		log.Println(err)
	}

	go func() {
		for message := range messages {
			log.Printf(" > Received message: %s\n", message.Body)

			data, err := services.ReadExcelResetPassword(string(message.Body))
			if err != nil {
				log.Println(err)
				return
			}

			for _, d := range data {
				body, _ := json.Marshal(d)
				config.Publish("reset-password", string(body))
			}
		}
	}()
}
