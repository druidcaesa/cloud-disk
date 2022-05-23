package models

import "time"

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

func (tableName User) TableName() string {
	return "user"
}
