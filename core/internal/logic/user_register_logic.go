package logic

import (
	"NetDisk/core/models"
	"NetDisk/core/utils"
	"context"
	"errors"
	"log"

	"NetDisk/core/internal/svc"
	"NetDisk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterReply, err error) {
	//判断code是否一致
	code, err := l.svcCtx.RDB.Get(l.ctx, req.Email).Result()
	if err != nil {
		return nil, errors.New("未获取该邮箱的验证码")
	} else if code != req.Code {
		err = errors.New("验证码错误")
		return
	}
	//判断用户名是否已存在
	var count int
	result := l.svcCtx.DB.Where("name = ?", req.Name).Table("user_basic").Count(&count)
	if result.Error != nil {
		return nil, nil
	} else if count > 0 {
		err = errors.New("用户名已存在")
		return
	}

	//数据入库
	userBasic := &models.UserBasic{
		Identity: utils.GetUuid(),
		Name:     req.Name,
		Password: utils.Md5(req.Password),
		Email:    req.Email,
	}
	result = l.svcCtx.DB.Create(userBasic)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Println("insert user row :", result.Row())
	return
}
