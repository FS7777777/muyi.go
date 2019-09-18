package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/muyi.go/bacc-utility/conf"
	"github.com/muyi.go/bacc-utility/protocol"
	"github.com/muyi.go/bacc-utility/util"
	"github.com/muyi.go/bacc-utility/ws"
	"net"
	"os"
)

func main() {
	fmt.Println(os.Args)
	r := gin.Default()
	r.LoadHTMLGlob("dist/index.html") // 添加入口index.html
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

	c := new(conf.Conf)
	c.GetConf()
	fmt.Println(c)

	// init web socket
	w := new(ws.WS)
	r.GET("/ping", w.Ping)
	//监听 get请求  /test路径
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, []string{"123", "321"})
		w.Send()
	})

	wg := util.WaitGroupWrapper{}
	// start TM server
	tcpListenerTM, err := net.Listen("tcp", ":8881")
	if err != nil {
		fmt.Println("listen error:", err)
		os.Exit(1)
	}
	wg.Wrap(func() {
		protocol.TCPTMServer(tcpListenerTM)
	})

	// start Image server
	tcpListenerImage, err := net.Listen("tcp", ":8882")
	if err != nil {
		fmt.Println("listen error:", err)
		os.Exit(1)
	}
	wg.Wrap(func() {
		protocol.TCPImageServer(tcpListenerImage)
	})

	// start tc loopback server
	tcpListenerTCLoopback, err := net.Listen("tcp", ":8883")
	if err != nil {
		fmt.Println("listen error:", err)
		os.Exit(1)
	}
	wg.Wrap(func() {
		protocol.TCPTCLoopbackServer(tcpListenerTCLoopback)
	})

	r.Run(":8080")
}
