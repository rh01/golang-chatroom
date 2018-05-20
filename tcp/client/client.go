package main

import (
	"fmt"
	"net"
	"log"
	"bufio"
	"os"
	"strings"
)

func main() {

	fmt.Println("start client ...")
	// start a client
	// 1. create a connection to server
	conn, e := net.Dial("tcp", "127.0.0.1:9000")
	if e != nil {
		log.Println(e)
		return
	}

	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for{

		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		if trimmedInput == "Q" {
			return
		}
		_, err := conn.Write([]byte(trimmedInput))
		if err!=nil {
			log.Println(err)
			return
		}
	}
}
