package logic

import (
	"NetDisk/core/models"
	"context"

	"NetDisk/core/internal/svc"
	"NetDisk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadPrepareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadPrepareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadPrepareLogic {
	return &FileUploadPrepareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadPrepareLogic) FileUploadPrepare(req *types.FileUploadPrepareRequest) (resp *types.FileUploadPrepareReply, err error) {
	rp := new(models.Repository_pool)
	result := l.svcCtx.DB.Where("hash = ?", req.Md5).Find(rp)
	if result.Error != nil {
		return nil, result.Error
	}
	resp = new(types.FileUploadPrepareReply)
	resp.Identity = rp.Identity
	return
}
