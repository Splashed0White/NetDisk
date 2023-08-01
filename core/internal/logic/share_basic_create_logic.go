package logic

import (
	"NetDisk/core/help"
	"NetDisk/core/internal/svc"
	"NetDisk/core/internal/types"
	"NetDisk/core/models"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicCreateLogic {
	return &ShareBasicCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicCreateLogic) ShareBasicCreate(req *types.ShareBasicCreateRequest, userIdentity string) (resp *types.ShareBasicCreateReply, err error) {
	// todo: add your logic here and delete this line
	uuid := help.GetUuid()
	ur := new(models.UserRepository)
	result := l.svcCtx.DB.Table("user_repository").Where("identity = ?", req.UserRepositoryIdentity).Find(ur)
	if result.Error != nil {
		return nil, result.Error
	}
	data := &models.ShareBasic{
		Identity:               uuid,
		UserIdentity:           userIdentity,
		UserRepositoryIdentity: req.UserRepositoryIdentity,
		RepositoryIdentity:     ur.RepositoryIdentity,
		ExpiredTime:            req.ExpiredTime,
	}
	result = l.svcCtx.DB.Create(data)
	if result.Error != nil {
		return nil, result.Error
	}
	resp = &types.ShareBasicCreateReply{Identity: uuid}
	return
}
