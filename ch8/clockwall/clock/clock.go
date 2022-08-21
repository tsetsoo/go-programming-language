package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

var loc *time.Location

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().In(loc).Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var err error
	loc, err = time.LoadLocation(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", os.Args[1]))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
