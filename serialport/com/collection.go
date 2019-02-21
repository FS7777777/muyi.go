package com

import (
	"fmt"
	"io"
	"log"
)

// Com采集传输对象
type ComCollection struct {
	reader      io.ReadWriteCloser
	comDataChan chan string
}

func New(iorwc io.ReadWriteCloser) *ComCollection {
	com := &ComCollection{
		reader:      iorwc,
		comDataChan: make(chan string),
	}
	return com
}

// 接收com口消息
func (c *ComCollection) receive() {
	for {
		buffer := make([]byte, 128)
		num, err := c.reader.Read(buffer)
		if err != nil {
			log.Fatal(err)
		}
		if num > 0 {
			str := string(buffer[:num])
			c.comDataChan <- str
		}
	}
}

//发送消息
func (c *ComCollection) send() {
	for {
		select {
		case v := <-c.comDataChan:
			fmt.Println("get channel")
			fmt.Println(string(v))
		}
	}
}

func (c *ComCollection) Main() {
	// 开启goroutine采集com口数据
	go c.receive()
	//发送消息
	go c.send()
}

// 关闭资源
func (c *ComCollection) Exit() {
	if c.reader != nil {
		c.reader.Close()
	}
	if c.comDataChan != nil {
		close(c.comDataChan)
	}
}
