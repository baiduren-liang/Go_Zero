package logic

import (
	"context"

	"GoZeroDemo/single/internal/svc"
	"GoZeroDemo/single/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SingleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSingleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SingleLogic {
	return &SingleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SingleLogic) Single(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return &types.Response{
		Message: "hello world",
	}, nil
}
