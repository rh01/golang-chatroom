package main

import (
	"github.com/garyburd/redigo/redis"
	"log"
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
	// Refer document:  https://redis.io/commands/
	// set "name" "shenhengheng"
	_, err = conn.Do("Set", "name", "shenhengheng")
	if err != nil {
		log.Println("set 命令错误，详细错误为：", err)
		return
	}

	// Get name
	reply, e := redis.String(conn.Do("GET", "name"))
	if e != nil {
		log.Println("不存在这样的key，详细错误：", e)
		return
	}

	fmt.Println(reply)
}
