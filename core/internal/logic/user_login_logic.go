package logic

import (
	"NetDisk/core/define"
	"NetDisk/core/models"
	"NetDisk/core/utils"
	"context"
	"fmt"

	"NetDisk/core/internal/svc"
	"NetDisk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginReply, err error) {
	// todo: add your logic here and delete this line
	//1.从数据库中查询用户
	userBasic := new(models.UserBasic)
	//fmt.Println(utils.Md5("123456"))
	result := l.svcCtx.DB.Where("name = ? AND password = ?", req.Name, utils.Md5(req.Password)).Find(userBasic)
	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}
	//2.生成token
	token, err := utils.GenerateToken(userBasic.Id, userBasic.Identity, userBasic.Name, 20)
	//token, err := utils.GenerateToken(1, "123456", "xiaoming")
	if err != nil {
		return nil, err
	}

	//3.用于刷新token的token
	refreshToken, err := utils.GenerateToken(userBasic.Id, userBasic.Identity, userBasic.Name, define.TokenExpire)
	resp = new(types.LoginReply)
	resp.Token = token
	resp.RefreshToken = refreshToken
	return
}
