package main

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

//func TestPingRoute(t *testing.T) {
//
//}

func main() {
	go broadcast1()
	http.ListenAndServe(":9090", nil)

}

func broadcast1()  {

	// 这里设置发送者的IP地址，自己查看一下自己的IP自行设定
	laddr := net.UDPAddr{
		//IP:   net.IPv4(192, 168, 123, 140),
		//Port: 3000,
	}

	// 这里设置接收者的IP地址为广播地址
	raddr := net.UDPAddr{
		IP:   net.IPv4(255, 255, 255, 255),
		//Port: 3000,
	}
	conn, err := net.DialUDP("udp", &laddr, &raddr)
	if err != nil {
		println(err.Error())
		return
	}

	for {
		_, err := conn.Write([]byte("hello world"))
		if err != nil{
			fmt.Println(err)
		}
		time.Sleep(time.Second)
	}
}