package article

import (
	"net/http"

	"GoZeroDemo/app/article/cmd/api/internal/logic/article"
	"GoZeroDemo/app/article/cmd/api/internal/svc"
	"GoZeroDemo/app/article/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获得文章列表
func GetArticlesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetArticleListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := article.NewGetArticlesLogic(r.Context(), svcCtx)
		resp, err := l.GetArticles(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
