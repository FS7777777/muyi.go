package main

import (
	"github.com/streadway/amqp"
	"log"
	"muyi.go/serialport/conf"
)

//定义MQ结构体
type MQ struct {
	amqp         string
	exchangeName string
	connection   *amqp.Connection
	ch           *amqp.Channel
}

//通过配置文件初始化MQ
func New(config conf.YamlConfig) *MQ {
	mq := &MQ{
		amqp:         config.Rabbit.Amqp,
		exchangeName: config.Rabbit.Exchange,
	}
	return mq
}

//定义统一异常处理
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

//初始化连接，定义通道、交换机
func (m *MQ) init() {
	conn, err := amqp.Dial(m.amqp)
	failOnError(err, "Failed to connect to RabbitMQ")
	m.connection = conn
	ch, err := conn.Channel()
	m.ch = ch
	failOnError(err, "Failed to open a channel")
	//定义Exchange
	exchangeDeclareErr := ch.ExchangeDeclare(
		m.exchangeName,
		"fanout",
		false,
		false,
		false,
		false,
		nil,
	)
	if exchangeDeclareErr != nil {
		failOnError(exchangeDeclareErr, "Failed to declare a exchange")
	}
}

//实现com DataTransfer接口
func (m MQ) Transfer(data []byte) {
	//发送消息
	body := "Hello World!"
	err := m.ch.Publish(
		m.exchangeName, // exchange
		"",             // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
}
