package modbus

import (
	"fmt"
	"github.com/goburrow/modbus"
	"time"
)

func Start() {
	// Modbus RTU/ASCII
	handler := modbus.NewRTUClientHandler("COM2")
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = "N"
	handler.StopBits = 1
	handler.SlaveId = 1
	handler.Timeout = 5 * time.Second

	err := handler.Connect()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer handler.Close()

	client := modbus.NewClient(handler)
	results, err := client.ReadHoldingRegisters(0, 10)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%d \n", results)
}

