package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	//addr, err := net.ResolveTCPAddr("tcp6", "[::]:0")
	addr, err := net.ResolveTCPAddr("tcp4", "0.0.0.0:0")
	if err != nil {
		log.Fatalln(err)
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("start success: " + listener.Addr().String())

	for {
		c, err := listener.AcceptTCP()
		if err != nil {
			fmt.Printf("accrpt error: %v\n", err)
			continue
		}
		go func(conn *net.TCPConn) {
			defer func() { _ = conn.Close() }()
			remoteAddr := conn.RemoteAddr().String()
			fmt.Printf("new tcp: %s\n", remoteAddr)
			_, _ = conn.Write([]byte(remoteAddr))
		}(c)
	}
}
