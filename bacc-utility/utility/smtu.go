package utility

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net"
	"net/http"
	"os"
	"time"
)

type SMTU struct {
	// log start time
	startTime time.Time
	// http server for web
	httpListener net.Listener
	// telemetry message
	tcpListenerTM net.Listener

	clientTM map[string]net.Conn
	// image
	tcpListenerImage net.Listener
	clientImage      map[string]net.Conn
	// telecontrol loopback
	tcpListenerTCLoopback net.Listener
	clientTCLoopback      map[string]net.Conn
	// log telecontrol
	tcpListenerTC net.Listener
	// log voice
	tcpListenerVoice net.Listener

	// ws
	ws *websocket.Conn
	// smtu Config
	smtuConfig *SMTUConfig
}

func Main() {
	fmt.Println(os.Args)
	smtu := &SMTU{
		startTime: time.Now(),
		// load smtu config info
		smtuConfig: (&SMTUConfig{}).GetConf(),

		clientTM:         make(map[string]net.Conn),
		clientImage:      make(map[string]net.Conn),
		clientTCLoopback: make(map[string]net.Conn),
	}
	fmt.Println(smtu)
	// start tcp
	smtu.initTCPServer()
	// start http server
	smtu.initHttpServer()
}

// init tcp server
func (smtu *SMTU) initTCPServer() {
	var err error
	// log tcp port
	fmt.Println(smtu.smtuConfig)
	wg := WaitGroupWrapper{}
	// start TM server
	smtu.tcpListenerTM, err = net.Listen("tcp", ":"+smtu.smtuConfig.TMPort)
	if err != nil {
		fmt.Println("listen error:", err)
		os.Exit(1)
	}
	wg.Wrap(func() {
		TCPTMServer(smtu.tcpListenerTM, smtu)
	})

	// start Image server
	smtu.tcpListenerImage, err = net.Listen("tcp", ":"+smtu.smtuConfig.ImagePort)
	if err != nil {
		fmt.Println("listen error:", err)
		os.Exit(1)
	}
	wg.Wrap(func() {
		TCPImageServer(smtu.tcpListenerImage, smtu)
	})

	// start tc loopback server
	smtu.tcpListenerTCLoopback, err = net.Listen("tcp", ":"+smtu.smtuConfig.TCLoopbackPort)
	if err != nil {
		fmt.Println("listen error:", err)
		os.Exit(1)
	}
	wg.Wrap(func() {
		TCPTCLoopbackServer(smtu.tcpListenerTCLoopback, smtu)
	})

	// start tc server
	smtu.tcpListenerTC, err = net.Listen("tcp", ":"+smtu.smtuConfig.TCPort)
	if err != nil {
		fmt.Println("listen error:", err)
		os.Exit(1)
	}
	wg.Wrap(func() {
		TCPTCServer(smtu.tcpListenerTC)
	})

	// start voice server
	smtu.tcpListenerVoice, err = net.Listen("tcp", ":"+smtu.smtuConfig.VoicePort)
	if err != nil {
		fmt.Println("listen error:", err)
		os.Exit(1)
	}
	wg.Wrap(func() {
		TCPVoiceServer(smtu.tcpListenerVoice)
	})
}

// add tcp client
func (smtu *SMTU) AddClientTM(clientID string, c net.Conn) {
	smtu.clientTM[clientID] = c
}
func (smtu *SMTU) AddClientImage(clientID string, c net.Conn) {
	smtu.clientImage[clientID] = c
}
func (smtu *SMTU) AddClientTCLoopback(clientID string, c net.Conn) {
	smtu.clientTCLoopback[clientID] = c
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

// ws server need http server to init
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (smtu *SMTU) ping(c *gin.Context) {
	var err error
	//升级get请求为webSocket协议
	smtu.ws, err = upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	err = smtu.ws.WriteMessage(1, []byte("hello we connected"))
	if err != nil {
		fmt.Println(err.Error())
	}
}
func (smtu *SMTU) send() {
	if smtu.ws == nil {
		return
	}
	for {
		//写入ws数据
		err := smtu.ws.WriteMessage(1, []byte("hello world"))
		if err != nil {
			break
		}
		time.Sleep(time.Duration(10) * time.Second)
	}
}

func (smtu *SMTU) initHttpServer() {
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
		v1.GET("/ping", smtu.ping)
		//监听 get请求  /test路径
		v1.GET("/test", func(c *gin.Context) {
			c.JSON(200, []string{"123", "321"})
			smtu.send()
		})
		v1.POST("/tm", func(c *gin.Context) {
			c.JSON(200, []string{"123", "321"})
			for _, v := range smtu.clientTM {
				writer := bufio.NewWriter(v)
				writer.Write([]byte("hello world i got you\n"))
				writer.Flush()
			}
		})
	}
	// start http server
	r.Run(":" + smtu.smtuConfig.HttpPort)
}
