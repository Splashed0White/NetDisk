package logic

import (
	"NetDisk/core/internal/svc"
	"NetDisk/core/internal/types"
	"NetDisk/core/models"
	"context"
	"errors"
	"github.com/jinzhu/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileNameUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileNameUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileNameUpdateLogic {
	return &UserFileNameUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileNameUpdateLogic) UserFileNameUpdate(req *types.UserFileNameUpdateRequest, userIdentity string) (resp *types.UserFileNameUpdateReply, err error) {
	//判断当前名称在该层级下是否存在
	var cnt int
	result := l.svcCtx.DB.Table("user_repository").Where("name = ? AND parent_id = (SELECT parent_id FROM user_repository ur WHERE ur.identity = ?)", req.Name, req.Identity).Count(&cnt)
	if result.Error != nil {
		return nil, result.Error
	}
	if cnt > 0 {
		return nil, errors.New("文件名称已存在")
	}
	//文件名称修改
	data := &models.UserRepository{
		Name: req.Name,
	}
	result = l.svcCtx.DB.Table("user_repository").Where("identity = ? AND user_identity = ?", req.Identity, userIdentity).Update(data)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("没查询到该记录")
		}
	}
	return
}
