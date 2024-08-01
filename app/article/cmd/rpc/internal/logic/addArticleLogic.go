package logic

import (
	"GoZeroDemo/app/article/model"
	"context"
	"fmt"

	"GoZeroDemo/app/article/cmd/rpc/internal/svc"
	"GoZeroDemo/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddArticleLogic {
	return &AddArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddArticle -----------------------article-----------------------
func (l *AddArticleLogic) AddArticle(in *pb.AddArticleReq) (*pb.AddArticleResp, error) {
	article := new(model.Article)
	article.Title = in.Title
	article.Content = in.Content
	// check
	res, _ := l.svcCtx.ArticleModel.FindByTitle(l.ctx, in.Title)
	fmt.Printf("res:%+v\n", res)
	_, err := l.svcCtx.ArticleModel.Insert(l.ctx, article)
	if err != nil {
		return nil, err
	}
	return &pb.AddArticleResp{}, nil

}
