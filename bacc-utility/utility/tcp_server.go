package utility

import (
	"fmt"
	"net"
	"os"
)

type tcpServer struct {
	ctx *smtuContext
}

func newTCPServer(ctx *smtuContext) {

	tcp := &tcpServer{
		ctx: ctx,
	}
	var err error
	// log tcp port
	fmt.Println(ctx.smtu.smtuConfig)
	//wg := WaitGroupWrapper{}
	// start TM server
	tcp.ctx.smtu.tcpListenerTM, err = net.Listen("tcp", ":"+tcp.ctx.smtu.smtuConfig.TMPort)
	if err != nil {
		fmt.Println("listen error:", err)
		os.Exit(1)
	}
	tcp.ctx.smtu.waitGroup.Wrap(func() {
		tcp.TCPTMServer(tcp.ctx.smtu.tcpListenerTM)
	})

	// start Image server
	tcp.ctx.smtu.tcpListenerImage, err = net.Listen("tcp", ":"+tcp.ctx.smtu.smtuConfig.ImagePort)
	if err != nil {
		fmt.Println("listen error:", err)
		os.Exit(1)
	}
	tcp.ctx.smtu.waitGroup.Wrap(func() {
		tcp.TCPImageServer(tcp.ctx.smtu.tcpListenerImage)
	})

	// start tc loopback server
	tcp.ctx.smtu.tcpListenerTCLoopback, err = net.Listen("tcp", ":"+tcp.ctx.smtu.smtuConfig.TCLoopbackPort)
	if err != nil {
		fmt.Println("listen error:", err)
		os.Exit(1)
	}
	tcp.ctx.smtu.waitGroup.Wrap(func() {
		tcp.TCPTCLoopbackServer(tcp.ctx.smtu.tcpListenerTCLoopback)
	})

	// start tc server
	tcp.ctx.smtu.tcpListenerTC, err = net.Listen("tcp", ":"+tcp.ctx.smtu.smtuConfig.TCPort)
	if err != nil {
		fmt.Println("listen error:", err)
		os.Exit(1)
	}
	tcp.ctx.smtu.waitGroup.Wrap(func() {
		tcp.TCPTCServer(tcp.ctx.smtu.tcpListenerTC)
	})

	// start voice server
	tcp.ctx.smtu.tcpListenerVoice, err = net.Listen("tcp", ":"+tcp.ctx.smtu.smtuConfig.VoicePort)
	if err != nil {
		fmt.Println("listen error:", err)
		os.Exit(1)
	}
	tcp.ctx.smtu.waitGroup.Wrap(func() {
		tcp.TCPVoiceServer(tcp.ctx.smtu.tcpListenerVoice)
	})
}
