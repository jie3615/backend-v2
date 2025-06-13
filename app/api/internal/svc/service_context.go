package svc

import (
	"metaLand/app/api/internal/config"
	"metaLand/app/api/internal/middleware"

	"github.com/redis/go-redis/v9"
	"github.com/sony/sonyflake"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config                       config.Config
	OIDCAuthMiddleware           rest.Middleware
	GuestAuthorizationMiddleware rest.Middleware
	DB                           *gorm.DB
	SF                           *sonyflake.Sonyflake
	RedisClient                  *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.DB.DataSource))
	if err != nil {
		panic(err)
	}
	sf := sonyflake.NewSonyflake(sonyflake.Settings{})
	if sf == nil {
		panic("sonyflake not created")
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Host,
		Password: c.Redis.Password,
		DB:       c.Redis.DB,
	})
	return &ServiceContext{
		Config:                       c,
		OIDCAuthMiddleware:           middleware.NewOIDCAuthMiddleware(c, db).Handle,
		GuestAuthorizationMiddleware: middleware.NewGuestAuthorizationMiddleware().Handle,
		DB:                           db,
		SF:                           sf,
		RedisClient:                  redisClient,
	}
}
