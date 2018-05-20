package main

import (
	"fmt"
	"net"
	"log"
)

func main() {

	fmt.Println("start server....")

	// listen address and port
	listener, e := net.Listen("tcp", "127.0.0.1:9000")
	if e != nil {
		log.Println(e)
		return
	}

	for {
		// Accept connection, return Conn and error
		conn, err := listener.Accept()
		if err != nil {
			log.Println("accept failed, error", err)
			continue
		}
		// process connection
		go process(conn)
	}

}
func process(conn net.Conn) {
	// 处理完就关闭
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			log.Println(err)
			return
		}
		// 就是简单的将客户端传过来的数据显示出来
		fmt.Println("read: ", string(buf))
	}
}
