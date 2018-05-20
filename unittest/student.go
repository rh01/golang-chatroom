package main

import (
	"encoding/json"
	"log"
	"io/ioutil"
)

type Student struct {
	Name string
	Age  int
	Male string
}

func (s *Student) Save() (err error) {
	// 序列化为json格式
	bytes, e := json.Marshal(s)
	if e != nil {
		log.Println(e)
		return
	}

	err = ioutil.WriteFile("d:/user.dat", bytes, 0755)
	return
}

func (s *Student) Load() (err error) {
	// 序列化为json格式
	bytes, err := ioutil.ReadFile("d:/user.dat")
	if err!= nil {
		log.Println(err)
		return 		
	}

	err = json.Unmarshal(bytes, s)
	return

}


