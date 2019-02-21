package main

import (
	"fmt"
	"github.com/judwhite/go-svc/svc"
	"github.com/tarm/goserial"
	"log"
	"muyi.go/serialport/com"
	"os"
	"path/filepath"
	"syscall"
)

// 使用go-svc进行管理程序初始化、启动、销毁工作

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
	if env.IsWindowsService() {
		dir := filepath.Dir(os.Args[0])
		return os.Chdir(dir)
	}
	//configInit()
	return nil
}

func (p *program) Start() error {
	// 连接配置
	conf := &serial.Config{Name: "COM2", Baud: 9600 /*毫秒*/ }
	iorwc, err := serial.OpenPort(conf)
	if err != nil {
		log.Fatal(err)
	}
	requestCommand := []byte{0x01, 0x03, 0x00, 0x00, 0x00, 0x02}
	collection := com.New(iorwc, requestCommand)

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
