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
	"strconv"
	"time"
)

type httpServer struct {
	ctx *smtuContext
}

func newHTTPServer(ctx *smtuContext) {
	s := &httpServer{
		ctx: ctx,
	}
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
		v1.GET("/ping", s.ping)
		//监听 get请求  /test路径
		v1.GET("/test", s.test)
		v1.POST("/upload_tm", s.uploadTM)
		v1.POST("/tm", s.tm)
		v1.POST("/upload_image", s.uploadImage)
		v1.POST("/image", s.image)
		v1.POST("/upload_tc", s.uploadTC)

		v1.POST("/tc", s.tc)

		v1.GET("/video", s.video)
	}
	// start http server
	r.Run(":" + s.ctx.smtu.smtuConfig.HttpPort)

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

func (s *httpServer) ping(c *gin.Context) {
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	s.ctx.smtu.clientSocket[ws.RemoteAddr().String()] = ws
	err = ws.WriteMessage(1, []byte("hello we connected"+ws.RemoteAddr().String()))
	if err != nil {
		fmt.Println(err.Error())
	}
}

// test send websocket
func (s *httpServer) test(c *gin.Context) {
	c.JSON(200, []string{"123", "321"})
	s.ctx.smtu.send()
}

// 上传遥测文件
func (s *httpServer) uploadTM(c *gin.Context) {
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
	s.ctx.smtu.dataTM = data
	// 上传文件到指定的路径
	// c.SaveUploadedFile(file, dst)
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

// 模拟遥测发送
func (s *httpServer) tm(c *gin.Context) {
	c.JSON(200, []string{"123", "321"})
	for _, v := range s.ctx.smtu.clientTM {
		writer := bufio.NewWriter(v)
		writer.Write(s.ctx.smtu.dataTM)
		writer.Flush()
	}
}

// 上传图片
func (s *httpServer) uploadImage(c *gin.Context) {
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
	s.ctx.smtu.dataImage = data
	// 上传文件到指定的路径
	// c.SaveUploadedFile(file, dst)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

// simulation send image
func (s *httpServer) image(c *gin.Context) {
	c.JSON(200, []string{"123", "321"})
	for _, v := range s.ctx.smtu.clientImage {
		writer := bufio.NewWriter(v)
		writer.Write(s.ctx.smtu.dataImage)
		writer.Flush()
	}
}

// upload tc
func (s *httpServer) uploadTC(c *gin.Context) {
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
	s.ctx.smtu.dataTC = data
	// 上传文件到指定的路径
	// c.SaveUploadedFile(file, dst)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

type TCPack struct {
	Ip     string `json:"ip"`
	Port   int    `json:"port"`
	Manual bool   `json:"manual" `
	First  string `json:"first" `
	Second string `json:"second" `
	Third  string `json:"third"`
}

// simulation send tc
func (s *httpServer) tc(c *gin.Context) {
	var tc TCPack

	err := c.ShouldBindJSON(&tc)
	if err != nil {
		log.Println(err)
	}
	log.Println(tc.Ip)

	// send udp
	addr, err := net.ResolveUDPAddr("udp", tc.Ip+":"+strconv.Itoa(tc.Port))
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

	_, err = conn.Write(s.ctx.smtu.dataTC)
	if err != nil {
		fmt.Println("failed:", err)
	}

	c.JSON(200, []string{"123", "321"})
	// udp 发送
}

// test play video
func (s *httpServer) video(c *gin.Context) {

	vl := "H:\\迅雷下载\\阳光电影www.ygdy8.com.大人物.HD.1080p.国语中英双字.mkv"

	video, err := os.Open(vl)
	if err != nil {
		log.Printf("Error when try to open file: %v", err)
		return
	}

	c.Header("Content-Type", "video/mp4")
	http.ServeContent(c.Writer, c.Request, "", time.Now(), video)

	defer video.Close()
}
