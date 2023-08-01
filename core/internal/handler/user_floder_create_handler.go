package handler

import (
	"net/http"

	"NetDisk/core/internal/logic"
	"NetDisk/core/internal/svc"
	"NetDisk/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserFloderCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserFloderCreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserFloderCreateLogic(r.Context(), svcCtx)
		resp, err := l.UserFloderCreate(&req, r.Header.Get("UserIdentity"))
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
