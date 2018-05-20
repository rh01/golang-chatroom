package proto

import . "learngo/chatroom/chat_server/model"

// 根据Cmd来选择Data，其中有两种Cmd格式
// 一是LoginCmd，接受id和密码
// 二是RegisterCmd, 接受User
type Message struct {
	Cmd string `json:"cmd"`
	Data string `json:"data"`
}

type LoginCmd struct {
	Id int `json:"id"`
	Passwd string `json:"passwd"`
}

type RegisterCmd struct {
	User User `json:"user"`
}

type LoginCmdRes struct {
	Code int `json:"code"`
	Error string `json:"error"`
}