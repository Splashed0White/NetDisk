package logic

import (
	"NetDisk/core/internal/svc"
	"NetDisk/core/internal/types"
	"NetDisk/core/models"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileMoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileMoveLogic {
	return &UserFileMoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileMoveLogic) UserFileMove(req *types.UserFileMoveRequest, userIdentity string) (resp *types.UserFileMoveReply, err error) {
	// todo: add your logic here and delete this line
	//查询文件夹是否存在
	parentData := new(models.UserRepository)
	result := l.svcCtx.DB.Where("identity = ? AND user_identity = ?", req.ParentIdentity, userIdentity).Find(parentData)
	if result.Error != nil {
		if result.RecordNotFound() {
			return nil, errors.New("文件夹不存在")
		}
		return
	}
	//更新记录的parentID

	// 这里如果删除了也会被放入到文件夹里面
	result = l.svcCtx.DB.Table("user_repository").Where("identity = ?", req.Identity).Update(models.UserRepository{
		ParentId: int64(parentData.Id),
	})
	if result.Error != nil {
		return nil, result.Error
	}
	return
}
