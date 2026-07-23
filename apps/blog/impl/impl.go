package impl

import (
	"blog/apps/blog"
	"blog/conf"
	"blog/ioc"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func init() {
	ioc.Contorller.Registry(blog.AppName, &BlogServiceImpl{})
}

type BlogServiceImpl struct {
	cache *redis.Client
	db *gorm.DB
}

func (i *BlogServiceImpl) Init() error {
	i.db = conf.C().Mysql.GetDB()
	i.cache=conf.C().Redis.GetClient()
	return nil
}
