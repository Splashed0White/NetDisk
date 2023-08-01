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

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *types.UserDetailRequest) (resp *types.UserDetailReply, err error) {
	// todo: add your logic here and delete this line
	resp = &types.UserDetailReply{}
	userBasic := new(models.UserBasic)
	result := l.svcCtx.DB.Where("identity = ?", req.Identity).Find(userBasic)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 处理记录不存在的情况
			return nil, errors.New("查找用户不存在")
		} else {
			return nil, result.Error
		}
	}
	resp.Name = userBasic.Name
	resp.Email = userBasic.Email
	return
}
