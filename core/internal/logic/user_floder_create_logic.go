package logic

import (
	"NetDisk/core/help"
	"NetDisk/core/internal/svc"
	"NetDisk/core/internal/types"
	"NetDisk/core/models"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserFloderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFloderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFloderCreateLogic {
	return &UserFloderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFloderCreateLogic) UserFloderCreate(req *types.UserFloderCreateRequest, userIdentity string) (resp *types.UserFloderCreateReply, err error) {
	// todo: add your logic here and delete this line
	//判断当前名称下该层级是否存在
	var cnt int
	result := l.svcCtx.DB.Table("user_repository").Where("name = ? AND parent_id = ?", req.Name, req.ParentId).Count(&cnt)
	if result.Error != nil {
		return nil, result.Error
	}
	if cnt > 0 {
		return nil, errors.New("文件名称已存在")
	}

	//创建文件夹
	data := &models.UserRepository{
		Identity:     help.GetUuid(),
		UserIdentity: userIdentity,
		ParentId:     req.ParentId,
		Name:         req.Name,
	}
	result = l.svcCtx.DB.Create(data)
	if result.Error != nil {
		return
	}
	return &types.UserFloderCreateReply{Identity: data.Identity}, nil
}
