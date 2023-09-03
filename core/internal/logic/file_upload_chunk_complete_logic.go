package logic

import (
	"NetDisk/core/define"
	"NetDisk/core/models"
	"NetDisk/core/utils"
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"

	"NetDisk/core/internal/svc"
	"NetDisk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadChunkCompleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadChunkCompleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadChunkCompleteLogic {
	return &FileUploadChunkCompleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadChunkCompleteLogic) FileUploadChunkComplete(req *types.FileUploadChunkCompleteRequest) (resp *types.FileUploadChunkCompleteReply, err error) {
	co := make([]cos.Object, 0)
	for _, v := range req.CosObjects {
		co = append(co, cos.Object{
			ETag:       v.Etag,
			PartNumber: v.PartNumber,
		})
	}
	err = utils.CosPartUploadComplete(req.Key, req.UploadId, co)

	// 数据入库
	rp := &models.Repository_pool{
		Identity: utils.GetUuid(),
		Hash:     req.Md5,
		Name:     req.Name,
		Ext:      req.Ext,
		Size:     req.Size,
		Path:     define.Url + "/" + req.Key,
	}
	l.svcCtx.DB.Create(rp)
	return
}
