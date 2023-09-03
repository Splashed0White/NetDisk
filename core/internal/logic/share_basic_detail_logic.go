package logic

import (
	"context"
	"errors"

	"NetDisk/core/internal/svc"
	"NetDisk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicDetailLogic {
	return &ShareBasicDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicDetailLogic) ShareBasicDetail(req *types.ShareBasicDetailRequest) (resp *types.ShareBasicDetailReply, err error) {

	//是否可以用事务进行优化
	//对分享记录的点击次数进行+1
	result := l.svcCtx.DB.Exec("UPDATE share_basic SET click_num = click_num + 1 WHERE identity = ?", req.Identity)
	if result.Error != nil {
		if result.RecordNotFound() {
			return nil, errors.New("没找到，点击次数加1")
		}
		return nil, result.Error
	}
	//获取资源的详细信息
	resp = new(types.ShareBasicDetailReply)
	result = l.svcCtx.DB.Table("share_basic").
		Select("share_basic.repository_identity,user_repository.name,repository_pool.ext,repository_pool.size,repository_pool.path").
		Joins("LEFT JOIN repository_pool ON share_basic.repository_identity = repository_pool.identity").Where("share_basic.identity = ?", req.Identity).
		Joins("LEFT JOIN user_repository ON user_repository.identity = share_basic.user_repository_identity").
		Find(resp)
	if result.Error != nil {
		if result.RecordNotFound() {
			return nil, errors.New("没找到")
		}
		return nil, result.Error
	}
	return
}
