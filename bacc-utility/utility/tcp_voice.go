package utility

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"runtime"
	"strings"
)

func handleVoiceConn(c net.Conn) {
	defer c.Close()
	reader := bufio.NewReader(c)
	for {
		b := make([]byte,100)
		_, err:= reader.Read(b)
		if err!=nil || err == io.EOF {
			fmt.Println(err.Error())
			break
		}
		fmt.Println(string(b))
	}
}


func TCPVoiceServer(listener net.Listener) {
	fmt.Println("TCP: listening on %s", listener.Addr())

	for {
		c, err := listener.Accept()
		if err != nil {
			if nerr, ok := err.(net.Error); ok && nerr.Temporary() {
				fmt.Println("temporary Accept() failure - %s", err)
				runtime.Gosched()
				continue
			}
			// theres no direct way to detect this error because it is not exposed
			if !strings.Contains(err.Error(), "use of closed network connection") {
				fmt.Println("listener.Accept() - %s", err)
			}
			break
		}
		fmt.Println("accept c:", c)
		// start a new goroutine to handle
		// the new connection.
		go handleVoiceConn(c)
	}
}
