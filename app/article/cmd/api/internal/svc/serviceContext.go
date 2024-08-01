package svc

import (
	"GoZeroDemo/app/article/cmd/api/internal/config"
	"GoZeroDemo/app/article/cmd/rpc/article"
	"github.com/zeromicro/go-zero/zrpc"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

type ServiceContext struct {
	Config     config.Config
	ArticleRpc article.ArticleZrpcClient
	EtcdClient *clientv3.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   c.ArticleRpcClient.Endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:     c,
		ArticleRpc: article.NewArticleZrpcClient(zrpc.MustNewClient(c.ArticleRpcClient)),
		EtcdClient: client,
	}
}
