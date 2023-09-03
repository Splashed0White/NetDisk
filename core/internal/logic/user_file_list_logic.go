package logic

import (
	"NetDisk/core/define"
	"context"
	"errors"
	"github.com/jinzhu/gorm"

	"NetDisk/core/internal/svc"
	"NetDisk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileListLogic) UserFileList(req *types.UserFileListRequest, userIdentity string) (resp *types.UserFileListReply, err error) {
	uf := make([]*types.UserFile, 0)
	resp = &types.UserFileListReply{}
	var cnt int64
	//分页参数
	size := req.Size
	if size == 0 {
		size = define.PageSize
	}

	page := req.Page
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * size //偏移量
	//查询用户文件列表
	result := l.svcCtx.DB.Table("user_repository").
		Select("user_repository.id, user_repository.identity,"+
			"user_repository.repository_identity,user_repository.ext, user_repository.name, repository_pool.path, repository_pool.size").
		Joins("LEFT JOIN repository_pool ON user_repository.repository_identity = repository_pool.identity").
		Offset(offset).Limit(size).
		Where("parent_id = ? AND user_identity = ? AND user_repository.deleted_at IS NULL", req.Id, userIdentity).
		Find(&uf)
	if result.Error != nil {
		//没查到
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("未查询到该List")
		}
		return nil, err
	}
	//查询用户文件总数
	result.Table("user_repository").
		Where("parent_id = ? AND user_identity = ?", req.Id, userIdentity).
		Count(&cnt)
	if result.Error != nil {
		return nil, result.Error
	}

	resp.List = uf
	resp.Count = cnt
	return
}
