package utility

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net"
	"net/http"
	"time"
)

type SMTU struct {
	// log start time
	startTime time.Time
	// http server for web
	// we not really use it, use gin instead of native http for web server
	httpListener net.Listener
	//gin server
	http *http.Server
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

	waitGroup WaitGroupWrapper

	dataTM    []byte
	dataImage []byte
	dataTC    []byte
}

func New() *SMTU {
	return &SMTU{
		startTime: time.Now(),
		// load smtu config info
		smtuConfig: (&SMTUConfig{}).GetConf(),

		clientTM:         make(map[string]net.Conn),
		clientImage:      make(map[string]net.Conn),
		clientTCLoopback: make(map[string]net.Conn),
		clientSocket:     make(map[string]*websocket.Conn),
	}
}

func (smtu *SMTU) Main() {
	ctx := &smtuContext{
		smtu: smtu,
	}
	fmt.Println(smtu)
	// start tcp
	newTCPServer(ctx)
	// start http server
	newHTTPServer(ctx)
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

func (smtu *SMTU) Exit() {
	if smtu.tcpListenerTM != nil {
		smtu.tcpListenerTM.Close()
	}

	if smtu.tcpListenerImage != nil {
		smtu.tcpListenerImage.Close()
	}

	if smtu.tcpListenerTC != nil {
		smtu.tcpListenerTC.Close()
	}

	if smtu.tcpListenerTCLoopback != nil {
		smtu.tcpListenerTCLoopback.Close()
	}

	if smtu.tcpListenerVoice != nil {
		smtu.tcpListenerVoice.Close()
	}

	if smtu.http != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := smtu.http.Shutdown(ctx); err != nil {
			log.Fatal("Server Shutdown:", err)
		}
		fmt.Println("shutdown http server exit")
	}

	smtu.waitGroup.Wait()
	fmt.Println("shutdown complete")
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
