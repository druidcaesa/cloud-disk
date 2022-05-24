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

// GetFileNameByUser 根据用户查询用户文件是否存在
func (ur UserRepository) GetFileNameByUser(req *types.UserFileNameEditRequest, engine *xorm.Engine) int64 {
	count, err := engine.Where("name = ? and parent_id = (select parent_id from user_repository ur Where ur.identity = ?)", req.Name, req.Identity).Count(&ur)
	if err != nil {
		return 0
	}
	return count
}

// Edit 修改数据
func (ur UserRepository) Edit(engine *xorm.Engine) (int64, error) {
	return engine.Where("identity=? AND user_identity=?", ur.Identity, ur.UserIdentity).Update(ur)
}

// GetByName 根据名称查询数据
func (ur UserRepository) GetByName(engine *xorm.Engine) (*UserRepository, error) {
	_, err := engine.Where("name = ? and parent_id = ?", ur.Name, ur.ParentId).Get(&ur)
	if err != nil {
		return nil, err
	}
	return &ur, nil
}

// Delete 删除用户文件数据
func (ur UserRepository) Delete(engine *xorm.Engine) (int64, error) {
	return engine.Where("identity = ? AND user_identity = ?", ur.Identity, ur.UserIdentity).Delete(&ur)
}

// GetByIdentityAndUserIdentity 根据Identity和UserIdentity查询资源
func (ur UserRepository) GetByIdentityAndUserIdentity(engine *xorm.Engine) (*UserRepository, error) {
	_, err := engine.Where("identity=? AND user_identity = ?", ur.Identity, ur.UserIdentity).Get(&ur)
	if err != nil {
		return nil, err
	}
	return &ur, nil
}
