package models

import (
	"time"
	"xorm.io/xorm"
)

type RepositoryPool struct {
	Id         int
	Identity   string
	Hash       string
	Name       string
	Ext        string
	Size       int64
	Path       string
	CreateTime time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"updated"`
	DeleteTime time.Time `xorm:"deleted"`
}

func (p RepositoryPool) TableName() string {
	return "repository_pool"
}

// GetHashByRepositoryPool 根据hash查询文件存储池中是否存在相同的文件
func (p RepositoryPool) GetHashByRepositoryPool(hash string, sql *xorm.Engine) (bool, error) {
	return sql.Where("hash=?", hash).Get(&p)
}

// Insert 保存文件存储池数据
func (p RepositoryPool) Insert(sql *xorm.Engine) (int64, error) {
	return sql.Insert(&p)
}
