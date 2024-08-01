package handler

import (
	"net/http"

	"GoZeroDemo/single/internal/logic"
	"GoZeroDemo/single/internal/svc"
	"GoZeroDemo/single/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PostDemoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PostSingleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewPostDemoLogic(r.Context(), svcCtx)
		resp, err := l.PostDemo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
