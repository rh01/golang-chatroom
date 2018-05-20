package model

// 定义状态常量, 枚举
const (
	UserStatusOnline = 1
	UserStatusOffline = iota
)

type User struct {
	UserId    int    `json:"user_id"`    // 用户id
	Passwd    string `json:"passwd"`     //用户密码
	Nick      string `json:"nick"`       //用户昵称
	Sex       string `json:"sex"`        //用户性别
	Header    string `json:"header"`     //用户头像
	LastLogin string `json:"last_login"` //最近一次登录
	Status    int    `json:"status"`     //用户状态
}


