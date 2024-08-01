package article

import (
	"GoZeroDemo/app/article/cmd/rpc/article"
	"context"

	"GoZeroDemo/app/article/cmd/api/internal/svc"
	"GoZeroDemo/app/article/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewCreateArticleLogic 创建文章
func NewCreateArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateArticleLogic {
	return &CreateArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateArticleLogic) CreateArticle(req *types.CreateArticleReq) (resp *types.CreateArticleResp, err error) {
	_, err = l.svcCtx.ArticleRpc.AddArticle(l.ctx, &article.AddArticleReq{
		Content: req.Content,
		Title:   req.Title,
	})
	if err != nil {
		return nil, err
	}
	return
}
