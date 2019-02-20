package main

import (
	"fmt"
	"github.com/tarm/goserial"
	"log"
	"time"
)

func main() {
	c := &serial.Config{Name: "COM2", Baud: 9600, ReadTimeout: time.Second * 5 /*毫秒*/ }

	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	//n, err := s.Write([]byte("test"))
	//if err != nil {
	//	log.Fatal(err)
	//}

	ch := make(chan []byte, 128)

	buf := make([]byte, 128)

	go func() {
		for {
			_, err := s.Read(buf)
			if err != nil {
				log.Fatal(err)
			}
			ch <- buf
		}
	}()

	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		}
	}
	defer s.Close()
}
