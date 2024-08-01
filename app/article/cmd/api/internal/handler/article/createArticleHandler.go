package article

import (
	"encoding/json"
	"fmt"
	"net/http"

	"GoZeroDemo/app/article/cmd/api/internal/logic/article"
	"GoZeroDemo/app/article/cmd/api/internal/svc"
	"GoZeroDemo/app/article/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// CreateArticleHandler 创建文章
func CreateArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateArticleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		b, err := json.Marshal(req)
		fmt.Printf("参数param:%s\n", b)
		l := article.NewCreateArticleLogic(r.Context(), svcCtx)
		resp, err := l.CreateArticle(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
