package logic

import (
	"NetDisk/core/help"
	"NetDisk/core/internal/svc"
	"NetDisk/core/internal/types"
	"NetDisk/core/models"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRepositorySaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRepositorySaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRepositorySaveLogic {
	return &UserRepositorySaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRepositorySaveLogic) UserRepositorySave(req *types.UserRepositorySaveRequest, UserIdentity string) (resp *types.UserRepositorySaveReply, err error) {
	// todo: add your logic here and delete this line
	ur := &models.UserRepository{
		Identity:           help.GetUuid(),
		UserIdentity:       UserIdentity,
		ParentId:           req.ParentId,
		RepositoryIdentity: req.RepositoryIdentity,
		Name:               req.Name,
		Ext:                req.Ext,
	}
	result := l.svcCtx.DB.Create(ur)
	if result.Error != nil {
		return nil, err
	}
	return &types.UserRepositorySaveReply{
		Identity: ur.Identity,
	}, nil
}
