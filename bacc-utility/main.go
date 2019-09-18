package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/muyi.go/bacc-utility/conf"
	"github.com/muyi.go/bacc-utility/protocol"
	"github.com/muyi.go/bacc-utility/util"
	"github.com/muyi.go/bacc-utility/ws"
	"net"
	"net/http"
	"os"
)

func main() {
	fmt.Println(os.Args)

	// start tcp
	initTCPServer()
	// start http server
	initHttpServer()
}

func initHttpServer()  {
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

	r.Use(cors())

	v1 := r.Group("/v1")
	{
		// init web socket
		v1.GET("/ping", ws.WebSocket.Ping)
		//监听 get请求  /test路径
		v1.GET("/test", func(c *gin.Context) {
			c.JSON(200, []string{"123", "321"})
			ws.WebSocket.Send()
		})
		v1.POST("/tm", func(c *gin.Context) {
			c.JSON(200, []string{"123", "321"})
		})
	}
	// start http server
	r.Run(":" + conf.PortConfig.HttpPort)
}

// init tcp server
func initTCPServer() {
	// log tcp port
	fmt.Println(conf.PortConfig)
	wg := util.WaitGroupWrapper{}
	// start TM server
	tcpListenerTM, err := net.Listen("tcp", ":"+conf.PortConfig.TMPort)
	if err != nil {
		fmt.Println("listen error:", err)
		os.Exit(1)
	}
	wg.Wrap(func() {
		protocol.TCPTMServer(tcpListenerTM)
	})

	// start Image server
	tcpListenerImage, err := net.Listen("tcp", ":"+conf.PortConfig.ImagePort)
	if err != nil {
		fmt.Println("listen error:", err)
		os.Exit(1)
	}
	wg.Wrap(func() {
		protocol.TCPImageServer(tcpListenerImage)
	})

	// start tc loopback server
	tcpListenerTCLoopback, err := net.Listen("tcp", ":"+conf.PortConfig.TCLoopbackPort)
	if err != nil {
		fmt.Println("listen error:", err)
		os.Exit(1)
	}
	wg.Wrap(func() {
		protocol.TCPTCLoopbackServer(tcpListenerTCLoopback)
	})

	// start tc server
	tcpListenerTC, err := net.Listen("tcp", ":"+conf.PortConfig.TCPort)
	if err != nil {
		fmt.Println("listen error:", err)
		os.Exit(1)
	}
	wg.Wrap(func() {
		protocol.TCPTCServer(tcpListenerTC)
	})

	// start voice server
	tcpListenerVoice, err := net.Listen("tcp", ":"+conf.PortConfig.VoicePort)
	if err != nil {
		fmt.Println("listen error:", err)
		os.Exit(1)
	}
	wg.Wrap(func() {
		protocol.TCPVoiceServer(tcpListenerVoice)
	})
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
