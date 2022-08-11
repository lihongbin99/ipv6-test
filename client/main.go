package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	localAddr, err := net.ResolveTCPAddr("tcp6", "[240e:3b2:3231:f9b0:61e6:ddec:159:5d43]:0")
	if err != nil {
		log.Fatalln(err)
	}
	serverAddr, err := net.ResolveTCPAddr("tcp6", "[2406:da18:3a:900:730d:ee3f:4864:496a]:13520")
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := net.DialTCP("tcp6", localAddr, serverAddr)
	if err != nil {
		log.Fatalln(err)
	}

	buf := make([]byte, 64*1024)
	readLen, err := conn.Read(buf)
	if err != nil && err != io.EOF {
		log.Fatalln(err)
	}

	fmt.Printf("addr: %s\n", string(buf[:readLen]))

	fmt.Println("exit")
	os.Stdin.Read([]byte{0})
}
