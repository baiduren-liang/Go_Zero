package article

import (
	"context"

	"GoZeroDemo/app/article/cmd/api/internal/svc"
	"GoZeroDemo/app/article/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticlesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获得文章列表
func NewGetArticlesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticlesLogic {
	return &GetArticlesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticlesLogic) GetArticles(req *types.GetArticleListReq) (resp *types.GetArticleListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
