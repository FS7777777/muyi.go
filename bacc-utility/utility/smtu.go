package utility

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
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
	clientSocket map[string]*websocket.Conn
	// smtu Config
	smtuConfig *SMTUConfig

	dataTM    []byte
	dataImage []byte
	dataTC    []byte
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
		clientSocket:     make(map[string]*websocket.Conn),
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
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	smtu.clientSocket[ws.RemoteAddr().String()] = ws
	err = ws.WriteMessage(1, []byte("hello we connected"+ws.RemoteAddr().String()))
	if err != nil {
		fmt.Println(err.Error())
	}
}
func (smtu *SMTU) send() {
	i := 0
	for {
		for k, v := range smtu.clientSocket {
			//写入ws数据
			err := v.WriteMessage(1, []byte(fmt.Sprintf("hello world %d", i)))
			if err != nil {
				fmt.Println(err.Error())
				delete(smtu.clientSocket, k)
				//continue
			}
		}

		i++
		time.Sleep(time.Duration(1) * time.Second)
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
		v1.POST("/upload_tm", func(c *gin.Context) {
			// 单文件
			file, _ := c.FormFile("file")
			log.Println(file.Filename)

			fileStream, err := file.Open()
			if err != nil {
				c.String(http.StatusInternalServerError, "read file error")
				return
			}
			data, err := ioutil.ReadAll(fileStream)
			if err != nil {
				c.String(http.StatusInternalServerError, "read file error")
				return
			}
			smtu.dataTM = data
			// 上传文件到指定的路径
			// c.SaveUploadedFile(file, dst)

			c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
		})
		v1.POST("/tm", func(c *gin.Context) {
			c.JSON(200, []string{"123", "321"})
			for _, v := range smtu.clientTM {
				writer := bufio.NewWriter(v)
				writer.Write(smtu.dataTM)
				writer.Flush()
			}
		})
		v1.POST("/upload_image", func(c *gin.Context) {
			// 单文件
			file, _ := c.FormFile("image")
			log.Println(file.Filename)

			fileStream, err := file.Open()
			if err != nil {
				c.String(http.StatusInternalServerError, "read file error")
				return
			}
			data, err := ioutil.ReadAll(fileStream)
			if err != nil {
				c.String(http.StatusInternalServerError, "read file error")
				return
			}
			smtu.dataImage = data
			// 上传文件到指定的路径
			// c.SaveUploadedFile(file, dst)

			c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
		})
		v1.POST("/image", func(c *gin.Context) {
			c.JSON(200, []string{"123", "321"})
			for _, v := range smtu.clientImage {
				writer := bufio.NewWriter(v)
				writer.Write(smtu.dataImage)
				writer.Flush()
			}
		})

		v1.POST("/upload_tc", func(c *gin.Context) {
			// 单文件
			file, _ := c.FormFile("file")
			log.Println(file.Filename)

			fileStream, err := file.Open()
			if err != nil {
				c.String(http.StatusInternalServerError, "read file error")
				return
			}
			data, err := ioutil.ReadAll(fileStream)
			if err != nil {
				c.String(http.StatusInternalServerError, "read file error")
				return
			}
			smtu.dataTC = data
			// 上传文件到指定的路径
			// c.SaveUploadedFile(file, dst)

			c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
		})

		type TCPack struct {
			Ip     string `json:"ip"`
			Port   int32  `json:"port"`
			Manual bool   `json:"manual" `
			First  string `json:"first" `
			Second string `json:"second" `
			Third  string `json:"third"`
		}

		v1.POST("/tc", func(c *gin.Context) {
			var tc TCPack

			err := c.ShouldBindJSON(&tc)
			if err != nil {
				log.Println(err)
			}
			log.Println(tc.Ip)

			// send udp
			addr, err := net.ResolveUDPAddr("udp", tc.Ip+":"+string(tc.Port))
			if err != nil {
				fmt.Println("Can't resolve address: ", err)
				c.JSON(200, []string{"Can't resolve address"})
			}

			conn, err := net.DialUDP("udp", nil, addr)
			if err != nil {
				fmt.Println("Can't dial: ", err)
				c.JSON(200, []string{"连接失败"})
			}
			defer conn.Close()

			_, err = conn.Write(smtu.dataTC)
			if err != nil {
				fmt.Println("failed:", err)
			}
			
			c.JSON(200, []string{"123", "321"})
			// udp 发送
		})
	}
	// start http server
	r.Run(":" + smtu.smtuConfig.HttpPort)
}
