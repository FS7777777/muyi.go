package utility

import (
	"bufio"
	"fmt"
	"net"
	"runtime"
	"strings"
	"time"
)

func handleImageConn(c net.Conn) {
	//defer c.Close()
	writer := bufio.NewWriter(c)
	//reader := bufio.NewReader(c)
	for {
		//data, err := ioutil.ReadFile("F:/6448355_204110337000_2.jpg")
		//if err != nil {
		//	fmt.Println("read error")
		//	return
		//}
		writer.Write([]byte("hello world\n"))
		writer.Flush()
		time.Sleep(time.Duration(3) * time.Second)
	}

}

func TCPImageServer(listener net.Listener, smtu *SMTU) {
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
		smtu.AddClientImage(c.RemoteAddr().String(), c)
		fmt.Println("accept c:", c)
		// start a new goroutine to handle
		// the new connection.
		//go handleImageConn(c)
	}
}
