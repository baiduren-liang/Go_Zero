package logic

import (
	"context"

	"GoZeroDemo/single/internal/svc"
	"GoZeroDemo/single/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostDemoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostDemoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostDemoLogic {
	return &PostDemoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostDemoLogic) PostDemo(req *types.PostSingleReq) (resp *types.PostSingleRes, err error) {
	return &types.PostSingleRes{
		Message: req.Message,
	}, nil
}
