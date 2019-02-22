package main

import (
	"github.com/streadway/amqp"
	"log"
)

//
func failOnReceive(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnReceive(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnReceive(err, "Failed to open a channel")
	defer ch.Close()

	//声明队列
	q, queueDeclareErr := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		true,    // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnReceive(queueDeclareErr, "Failed to declare a queue")

	//绑定队列到交换机
	queueBindErr := ch.QueueBind(q.Name, "", "fanoutComExchange", false, nil)
	failOnReceive(queueBindErr, "Failed to declare a queue")

	//订阅消息
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnReceive(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
