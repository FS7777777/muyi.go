package com

import (
	"fmt"
	"github.com/tarm/goserial"
	"io"
	"log"
	"time"
)

// 接收com口消息
func receive(ch chan<- string, Reader io.Reader) {
	for {
		buffer := make([]byte, 128)
		num, err := Reader.Read(buffer)
		if err != nil {
			log.Fatal(err)
		}
		if num > 0 {
			str := string(buffer[:num])
			ch <- str
		}
	}
}

//发送消息
func send(ch <-chan string) {
	for {
		select {
		case v := <-ch:
			fmt.Println("get channel")
			fmt.Println(string(v))
		}
	}
}

func Main() {
	// 连接配置
	c := &serial.Config{Name: "COM2", Baud: 9600, ReadTimeout: time.Second * 1 /*毫秒*/ }
	iorwc, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	//定义消息传输通道
	comDataChan := make(chan string)
	// 开启goroutine采集com口数据
	go receive(comDataChan, iorwc)
	//发送消息
	send(comDataChan)
	//关闭工作
	defer iorwc.Close()
	defer close(comDataChan)
}
