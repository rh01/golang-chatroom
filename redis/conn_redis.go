package main

import (

	"log"
	"github.com/garyburd/redigo/redis"
)

func main() {

	// start a connection to redis server
	// same as general client
	conn, e := redis.Dial("tcp", "124.16.87.91:6379")
	if e != nil {
		log.Println("连接失败, detail:", e)
		return
	}
	log.Println("连接成功...")
	defer conn.Close()



}
