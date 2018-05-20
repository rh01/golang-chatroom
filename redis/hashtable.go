package main

import (
	"log"
	"github.com/garyburd/redigo/redis"
	"fmt"
)

func main() {
	// connect to redis-server
	conn, err := redis.Dial("tcp", "124.16.87.91:6379")
	if err != nil {
		log.Println("连接redis server失败, 详细错误为：", err)
		return
	}

	defer conn.Close()
	// Refer document:  https://redis.io/command/hashtable
	_, err = conn.Do("HSet", "books", "abc", 100)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.Int(conn.Do("HGet", "books", "abc"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}

	fmt.Println(r)



}
