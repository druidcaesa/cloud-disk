package user

import (
	"cloud-disk/models"
	"cloud-disk/result"
	"cloud-disk/utils"
	"context"

	"cloud-disk/internal/svc"
	"cloud-disk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDeleteFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDeleteFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDeleteFileLogic {
	return &UserDeleteFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDeleteFileLogic) UserDeleteFile(req *types.UserDeleteFileRequest, userIdentity string) (resp *types.UserDeleteFileResponse, err error) {
	resp = &types.UserDeleteFileResponse{}
	//删除用户文件
	// TODO 这里后期应该会有下级文件或者文件夹的删除逻辑，目前先简单处理。实际逻辑请自行处理
	ur := new(models.UserRepository)
	ur.Identity = req.Identity
	ur.UserIdentity = userIdentity
	_, err = ur.Delete(l.svcCtx.Engine)
	if err != nil {
		resp.Result = result.ERROR(utils.FormatErrorLog(err))
		return resp, nil
	}
	resp.Result = result.OK()
	return
}
