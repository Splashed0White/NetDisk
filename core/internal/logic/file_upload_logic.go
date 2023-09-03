package logic

import (
	"NetDisk/core/models"
	"NetDisk/core/utils"
	"context"
	"time"

	"NetDisk/core/internal/svc"
	"NetDisk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadLogic) FileUpload(req *types.FileUploadRequest) (resp *types.FileUploadReply, err error) {

	//上传时没有做文件是否存在的判断
	rp := &models.Repository_pool{
		Identity:   utils.GetUuid(),
		Hash:       req.Hash,
		Name:       req.Name,
		Ext:        req.Ext,
		Size:       req.Size,
		Path:       req.Path,
		Created_at: time.Time{},
		Updated_at: time.Time{},
		Deleted_at: nil,
	}
	result := l.svcCtx.DB.Create(rp)
	if result.Error != nil {
		return
	}
	resp = new(types.FileUploadReply)
	resp.Identity = rp.Identity
	resp.Ext = rp.Ext
	resp.Name = rp.Name
	return
}
