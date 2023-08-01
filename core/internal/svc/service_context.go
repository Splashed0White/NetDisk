package svc

import (
	"NetDisk/core/internal/config"
	"NetDisk/core/internal/middleware"
	"NetDisk/core/models"
	"github.com/jinzhu/gorm"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	RDB    *redis.Client
	Auth   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     models.Init(c.MySQL.DataSource),
		RDB:    models.InitRedis(c.Redis.Addr),
		Auth:   middleware.NewAuthMiddleware().Handle,
	}
}
