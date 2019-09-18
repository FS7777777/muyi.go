package ws

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WS struct {
	Server *websocket.Conn
	err    error
}

var (
	WebSocket = &WS{}
)
//webSocket请求ping 返回pong
func (ws *WS) Ping(c *gin.Context) {
	//升级get请求为webSocket协议
	ws.Server, ws.err = upGrader.Upgrade(c.Writer, c.Request, nil)
	if ws.err != nil {
		return
	}
	err := ws.Server.WriteMessage(1, []byte("hello we connected"))
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (ws *WS) close() {
	ws.Server.Close()
}

func (ws *WS) Send() {
	if ws == nil {
		return
	}
	for {
		//读取ws中的数据
		//mt, message, err := ws.ReadMessage()
		//if err != nil {
		//	break
		//}
		//if string(message) == "ping" {
		//	message = []byte("pong")
		//}
		//写入ws数据
		err := ws.Server.WriteMessage(1, []byte("hello world"))
		if err != nil {
			break
		}
		time.Sleep(time.Duration(10) * time.Second)
	}
}
