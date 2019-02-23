package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

// 定义失败回调
func failOnSend(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// 声明fanout ExchangeName
var fanoutExchangeName = "fanoutComExchange"

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnSend(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnSend(err, "Failed to open a channel")
	defer ch.Close()

	//定义Exchange
	exchangeDeclareErr := ch.ExchangeDeclare(
		fanoutExchangeName,
		"fanout",
		false,
		false,
		false,
		false,
		nil,
	)
	if exchangeDeclareErr != nil {
		fmt.Println(err)
		panic(exchangeDeclareErr)
	}

	failOnSend(exchangeDeclareErr, "Failed to declare a exchange")

	//发送消息
	body := "Hello World!"
	err = ch.Publish(
		fanoutExchangeName, // exchange
		"",                 // routing key
		false,              // mandatory
		false,              // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	failOnSend(err, "Failed to publish a message")
}
