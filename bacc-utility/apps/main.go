package main

import (
	"fmt"
	"github.com/judwhite/go-svc/svc"
	"github.com/muyi.go/bacc-utility/utility"
	"log"
	"syscall"
)

type program struct {
	smtu *utility.SMTU
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
	smtu := utility.New()
	fmt.Println("start.....")
	fmt.Println(syscall.Getpid())
	go smtu.Main()
	p.smtu = smtu
	return nil
}
func (p *program) Stop() error {

	fmt.Println("stop.....")
	p.smtu.Exit()
	return nil
}
