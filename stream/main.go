package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"os"
)

func main(){
	fmt.Println(os.Args)
	r := gin.Default()
	r.LoadHTMLGlob("dist/index.html")    // 添加入口index.html
	//r.LoadHTMLFiles("static/*/*")	// 添加资源路径
	r.Static("/assets", "./dist/assets")
	r.Static("/css", "./dist/css")
	r.Static("/img", "./dist/img")
	r.Static("/js", "./dist/js")
	r.Static("/loading", "./dist/loading")
	r.Static("avatar2.jpg", "./dist/avatar2.jpg")
	r.Static("logo.png", "./dist/logo.png")
	r.Static("color.less", "./dist/color.less")
	r.StaticFile("/", "./dist/index.html")

	r.Run(":8080")
}

func checkError(err error){
	if  err != nil {
		fmt.Println("Error: %s", err.Error())
		os.Exit(1)
	}
}

func recvUDPMsg(conn *net.UDPConn){
	var buf [20]byte

	n, raddr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}

	fmt.Println("msg is ", string(buf[0:n]))

	//WriteToUDP
	//func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error)
	_, err = conn.WriteToUDP([]byte("nice to see u"), raddr)
	checkError(err)
}

func main() {
	udp_addr, err := net.ResolveUDPAddr("udp", ":11110")
	checkError(err)

	conn, err := net.ListenUDP("udp", udp_addr)
	defer conn.Close()
	checkError(err)

	//go recvUDPMsg(conn)
	recvUDPMsg(conn)
}