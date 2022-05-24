package models

import (
	"cloud-disk/define"
	"cloud-disk/internal/types"
	"time"
	"xorm.io/xorm"
)

type UserRepository struct {
	Id                 int
	Identity           string
	UserIdentity       string
	ParentId           int64
	RepositoryIdentity string
	Ext                string
	Name               string
	CreateTime         time.Time `xorm:"created"`
	UpdateTime         time.Time `xorm:"updated"`
	DeleteTime         time.Time `xorm:"deleted"`
}

func (ur UserRepository) TableName() string {
	return "user_repository"
}

// Insert 保存用户文件仓库池
func (ur UserRepository) Insert(engine *xorm.Engine) (int64, error) {
	return engine.Insert(&ur)
}

// UserFileList 根据用户查询用户文件列表
func (ur UserRepository) UserFileList(req *types.UserFileListRequest, userIdentity string, engine *xorm.Engine) ([]*types.UserFile, error) {
	uf := make([]*types.UserFile, 0)
	//分页参数
	size := req.Size
	if size == 0 {
		size = define.PageSize
	}
	page := req.Page
	if page == 0 {
		page = 1
	}
	err := engine.Table("user_repository").Where("parent_id=? and user_identity=?", req.Id, userIdentity).
		Select("user_repository.id,user_repository.identity,user_repository.repository_identity,user_repository.ext,"+
			"user_repository.name,repository_pool.path,repository_pool.size").
		Join("LEFT", "repository_pool", "user_repository.repository_identity=repository_pool.identity").
		Where("user_repository.delete_time = ? OR user_repository.delete_time IS NULL", time.Time{}.Format(define.DateTime)).
		Limit(size, (page-1)*size).Find(&uf)
	if err != nil {
		return nil, err
	}
	return uf, nil
}
