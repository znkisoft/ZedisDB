package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()
	log.Printf("[running] bound to %q", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("[error] %q", err.Error())
		}
		go func(c net.Conn) {
			handle(c)
		}(conn)

	}
}

func handle(conn net.Conn) {
	buf := make([]byte, 512)
	length, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading", err.Error())
		return
	}
	fmt.Printf("Received data: %v, length: %d", string(buf[:length]), length)
}
