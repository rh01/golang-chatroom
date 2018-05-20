package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
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
	_, err = conn.Do("MSet", "abc", 100, "efg", 300)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.Ints(conn.Do("MGet", "abc", "efg"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}

	for _, v := range r {
		fmt.Println(v)
	}

}
