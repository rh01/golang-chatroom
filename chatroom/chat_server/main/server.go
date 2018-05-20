package main

import (
	"net"
	"fmt"
)

func runServer(addr string) (err error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed, ", err)
			continue
		}
		go process(conn)
	}

}
func process(conn net.Conn) {

	defer conn.Close()
	client := &Client{
		conn: conn,
	}

	err := client.Process()
	if err != nil {
		fmt.Println("client process failed, ", err)
		return
	}
}
