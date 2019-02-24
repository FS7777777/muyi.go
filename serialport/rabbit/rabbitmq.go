package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"time"
)

// RabbitMQ 用于管理和维护rabbitmq的对象
type RabbitMQ struct {
	amqp         string
	connection   *amqp.Connection
	channel      *amqp.Channel
	exchangeName string
	exchangeType string // exchange的类型
}

// New 创建一个新的操作RabbitMQ的对象
func New() *RabbitMQ {
	// 这里可以根据自己的需要去定义
	return &RabbitMQ{
		amqp:         "amqp://guest:guest@localhost:5672/",
		exchangeName: "fanoutComExchange",
		exchangeType: "fanout",
	}
}

//声明MQ发送方所定义的方法
type sender interface {
	//连接
	conn() error
	//判断连接是否有效
	isClosed() bool
	//创建channel
	createChannel() error
	//定义交换机
	createExchange() error
	//刷新连接
	refresh()
	//消息发送
	publish()
	//启动
	Run()
	//关闭
	Close()
}

func (mq *RabbitMQ) conn() error {
	conn, err := amqp.Dial(mq.amqp)
	mq.connection = conn
	return err
}

func (mq *RabbitMQ) isClosed() bool {
	if mq.connection == nil {
		return true
	}
	return mq.connection.IsClosed()
}

func (mq *RabbitMQ) createChannel() error {
	ch, err := mq.connection.Channel()
	mq.channel = ch
	return err
}

// createExchange 准备rabbitmq的Exchange
func (mq *RabbitMQ) createExchange() error {
	// 申明Exchange
	err := mq.channel.ExchangeDeclare(
		mq.exchangeName, // exchange
		mq.exchangeType, // type
		false,           // durable
		false,           // autoDelete
		false,           // internal
		false,           // noWait
		nil,             // args
	)

	if nil != err {
		return err
	}

	return nil
}

func (mq *RabbitMQ) Close() {
	fmt.Println("close")

	if mq.channel != nil {
		mq.channel.Close()
		fmt.Println("close channel")
	}

	if mq.connection != nil {
		mq.connection.Close()
		fmt.Println("close connection")
	}
}

// run 开始获取连接并初始化相关操作
func (mq *RabbitMQ) refresh() {
	if !mq.isClosed() {
		fmt.Println("rabbit已连接")
		return
	}
	mq.Close()

	err := mq.conn()
	if err != nil {
		fmt.Println("rabbit刷新连接失败，将要重连:")
		return
	}

	// 获取新的channel对象
	mq.createChannel()

	// 初始化Exchange
	mq.createExchange()
}

// Start 启动Rabbitmq的客户端
func (mq *RabbitMQ) Run() {
	mq.refresh()
	go func() {
		ticker := time.NewTicker(2 * time.Second)
		for t := range ticker.C {
			fmt.Println("refresh mq conn", t)
			mq.refresh()
		}
	}()
}

func (mq *RabbitMQ) publish() {
	//发送消息
	body := "Hello World!"
	if mq.channel == nil {
		fmt.Println("channel close")
		return
	}
	err := mq.channel.Publish(
		mq.exchangeName, // exchange
		"",              // routing key
		false,           // mandatory
		false,           // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	fmt.Println(" [x] Sent %s", body)
	if err != nil {
		fmt.Println("channel publish error", err)
	}
}

func main() {
	forever := make(chan bool)
	mq := New()
	mq.Run()
	go func() {
		ticker := time.NewTicker(2 * time.Second)
		for t := range ticker.C {
			fmt.Println("publish", t)
			mq.publish()
		}
	}()
	<-forever
}
