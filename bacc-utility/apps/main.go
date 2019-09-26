package main

import (
	"fmt"
	"github.com/judwhite/go-svc/svc"
	"github.com/muyi.go/bacc-utility/utility"
	"log"
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
	go utility.Main()
	return nil
}
func (program) Stop() error {

	fmt.Println("stop.....")
	return nil
}