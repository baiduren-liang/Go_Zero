package svc

import (
	"GoZeroDemo/app/article/cmd/rpc/internal/config"
	"GoZeroDemo/app/article/model"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	redisClient  *redis.Redis
	ArticleModel model.ArticleModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		ArticleModel: model.NewArticleModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
