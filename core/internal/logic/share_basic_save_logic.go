package logic

import (
	"NetDisk/core/internal/svc"
	"NetDisk/core/internal/types"
	"NetDisk/core/models"
	"NetDisk/core/utils"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicSaveLogic {
	return &ShareBasicSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicSaveLogic) ShareBasicSave(req *types.ShareBasicSaveRequest, userIdentity string) (resp *types.ShareBasicSaveReply, err error) {
	//获取资源详情
	rp := new(models.Repository_pool)
	result := l.svcCtx.DB.Where("identity = ?", req.RepositoryIdentity).Find(rp)
	if result.Error != nil {
		if result.RecordNotFound() {
			return nil, errors.New("资源不存在")
		}
		return nil, result.Error
	}

	//UserReposity资源保存
	ur := &models.UserRepository{
		Identity:           utils.GetUuid(),
		UserIdentity:       userIdentity,
		ParentId:           req.ParentId,
		RepositoryIdentity: req.RepositoryIdentity,
		Name:               rp.Ext,
		Ext:                rp.Name,
	}
	result = l.svcCtx.DB.Create(ur)
	if result.Error != nil {
		return nil, result.Error
	}
	resp = new(types.ShareBasicSaveReply)
	resp.Identity = ur.Identity
	return
}
