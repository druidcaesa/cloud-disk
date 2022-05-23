package models

import (
	"errors"
	"time"
	"xorm.io/xorm"
)

type User struct {
	Id         int
	UserName   string
	Identity   string
	Password   string
	Email      string
	CreateTime time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"updated"`
	DeleteTime time.Time `xorm:"deleted"`
}

type UserInfo struct {
	UserName   string    `json:"userName"`
	Identity   string    `json:"identity"`
	Email      string    `json:"email"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	DeleteTime time.Time `json:"deleteTime"`
}

func (u User) TableName() string {
	return "user"
}

// GetUserInfo 获取当前用户详情
func (u User) GetUserInfo(identity string, engine *xorm.Engine) *UserInfo {
	userInfo := new(UserInfo)
	_, err := engine.Table(u.TableName()).Where("identity = ?", identity).Get(userInfo)
	if err != nil {
		return nil
	}
	return userInfo
}

// GetUserByUsername 根据用户名查询用户
func (u User) GetUserByUsername(username string, engine *xorm.Engine) (*User, error) {
	get, err := engine.Where("user_name = ?", username).Get(&u)
	if err != nil {
		return nil, err
	}
	if !get {
		return nil, errors.New("用户不存在")
	}
	return &u, nil
}
