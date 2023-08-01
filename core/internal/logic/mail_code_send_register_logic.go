package logic

import (
	"NetDisk/core/define"
	"NetDisk/core/help"
	"NetDisk/core/internal/svc"
	"NetDisk/core/internal/types"
	"context"
	"errors"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type MailCodeSendRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMailCodeSendRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MailCodeSendRegisterLogic {
	return &MailCodeSendRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MailCodeSendRegisterLogic) MailCodeSendRegister(req *types.MailCodeSendRequest) (resp *types.MailCodeSendReply, err error) {
	// todo: add your logic here and delete this line
	var count int
	result := l.svcCtx.DB.Where("email = ?", req.Email).Table("user_basic").Count(&count)
	if result.Error != nil {
		return
	} else {
		if count > 0 {
			err = errors.New("该邮箱已经被注册")
			return
		}
	}
	//该邮箱未被注册
	code := help.RandCode()
	//存储验证码
	l.svcCtx.RDB.Set(l.ctx, req.Email, code, time.Second*time.Duration(define.CodeExpire))
	//发送验证码
	err = help.MailCodeSend(req.Email, code)
	if err != nil {
		return nil, err
	}
	return nil, errors.New(req.Email)
}
