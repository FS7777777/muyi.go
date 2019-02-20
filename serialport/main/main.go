package main

import (
	"fmt"
	"github.com/judwhite/go-svc/svc"
	"log"
	"muyi.go/serialport/com"
	"syscall"
)

type program struct {
}

func main() {
	prg := &program{}
	if err := svc.Run(prg, syscall.SIGINT, syscall.SIGTERM); err != nil {
		log.Fatal(err)
	}
}

func (program) Init(env svc.Environment) error {

	fmt.Println("init.....")
	return nil
}

func (program) Start() error {

	fmt.Println("start.....")
	fmt.Println(syscall.Getpid())
	go func() {
		com.Main()
	}()
	return nil
}
func (program) Stop() error {

	fmt.Println("stop.....")
	return nil
}
