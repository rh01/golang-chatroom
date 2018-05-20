package model

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
	"encoding/json"
	"time"
)

// 用户管理
// redis table name of HashTable
var UserTable string = "users"

// redis 连接池
type UserMgr struct {
	poll *redis.Pool
}

// 创建一个新的用户管理
func NewUserMgr(pool *redis.Pool) (mgr *UserMgr) {
	mgr = &UserMgr{
		poll: pool,
	}
	return mgr
}

// 获取用户
func (p *UserMgr) getUser(conn redis.Conn, id int) (user *User, err error) {
	result, err := redis.String(
		conn.Do("HGet",
			UserTable,
			fmt.Sprintf("%d", id),
		))
	if err != nil {
		if err == redis.ErrNil {
			err = ErrUserNotExist
		}
		return
	}

	// 下面两个方式都可以
	//user := &User{}
	user = new(User)

	// 反序列化得到目标，然后赋值给user
	err = json.Unmarshal([]byte(result), user)
	if err != nil {
		return
	}
	return
}

// 用户登录
func (p *UserMgr) Login(id int, passwd string) (user *User, err error) {
	// 获得当前的连接
	conn := p.poll.Get()
	defer conn.Close()

	user, err = p.getUser(conn, id)
	if err != nil {
		return
	}

	if user.UserId != id || user.Passwd != passwd {
		err = ErrInvalidPasswd
	}

	user.Status = UserStatusOnline
	user.LastLogin = fmt.Sprintf("%v", time.Now())
	return
}

// 用户注册
func (p *UserMgr) Register(user *User) (err error) {
	// 获得当前连接
	conn := p.poll.Get()
	defer conn.Close()

	if user == nil {
		fmt.Println("invalid user")
		err = ErrInvalidParams
		return
	}

	// 通过传进来的用户id来查看当前数据库中是否存在此用户，
	// 若err为空，则说明存在此用具，返回ErrUserExist
	_, err = p.getUser(conn, user.UserId)
	if err == nil {
		err = ErrUserExist
		return
	}

	// 若发生其他错误，返回
	if err != ErrUserNotExist {
		return
	}

	data, err := json.Marshal(user)

	// HashTable 存储格式为  HSet "user" "id" "data"
	_, err = conn.Do("HSet", UserTable, fmt.Sprintf("%d", user.UserId), data)
	if err != nil {
		return
	}
	return
}


