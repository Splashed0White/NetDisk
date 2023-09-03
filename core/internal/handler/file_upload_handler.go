package handler

import (
	"NetDisk/core/models"
	"NetDisk/core/utils"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
	"path"

	"NetDisk/core/internal/logic"
	"NetDisk/core/internal/svc"
	"NetDisk/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			return
		}
		//判断文件是否已存在
		b := make([]byte, fileHeader.Size)
		_, err = file.Read(b)

		if err != nil {
			return
		}
		hash := fmt.Sprintf("%x", md5.Sum(b))
		rp := new(models.Repository_pool)
		result := svcCtx.DB.Where("hash = ?", hash).Find(rp)
		if result.Error != nil {
			//文件不存在,往cos中存
			//fmt.Println("文件不存在")
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				cosPath, err := utils.CosUpload(r)
				if err != nil {
					return
				}

				//往logic中传递req
				req.Name = fileHeader.Filename
				req.Ext = path.Ext(fileHeader.Filename)
				req.Size = fileHeader.Size
				req.Hash = hash
				req.Path = cosPath
			} else {
				return
			}
		} else { //文件存在
			//fmt.Println("文件存在")
			httpx.OkJson(w, &types.FileUploadReply{Identity: rp.Identity, Ext: rp.Ext, Name: rp.Name})
			return
		}

		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
