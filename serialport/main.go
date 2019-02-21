package main

import (
	"fmt"
	"github.com/judwhite/go-svc/svc"
	"github.com/tarm/goserial"
	"log"
	"muyi.go/serialport/com"
	"syscall"
	"time"
)

type program struct {
	collection *com.ComCollection
}

func main() {
	prg := &program{}
	if err := svc.Run(prg, syscall.SIGINT, syscall.SIGTERM); err != nil {
		log.Fatal(err)
	}
}

func (p *program) Init(env svc.Environment) error {
	fmt.Println("init.....")
	return nil
}

func (p *program) Start() error {
	// 连接配置
	conf := &serial.Config{Name: "COM2", Baud: 9600, ReadTimeout: time.Second * 1 /*毫秒*/ }
	iorwc, err := serial.OpenPort(conf)
	if err != nil {
		log.Fatal(err)
	}
	collection := com.New(iorwc)

	fmt.Println("start.....")
	fmt.Println(syscall.Getpid())

	collection.Main()
	p.collection = collection
	return nil
}
func (p *program) Stop() error {
	//关闭工作
	p.collection.Exit()
	fmt.Println("stop.....")
	return nil
}
