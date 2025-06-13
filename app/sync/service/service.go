package service

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"metaLand/app/sync/service/bounty"
	"metaLand/app/sync/service/common"
	"metaLand/app/sync/service/config"
	"metaLand/app/sync/service/crontask"
	"metaLand/app/sync/service/startup"
)

type Service struct {
	ctx *common.ServiceContext
}

func New(ctx context.Context, cfg *config.Config) (*Service, error) {
	db, err := gorm.Open(mysql.Open(cfg.DB.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	redis := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		DB:   cfg.Redis.DB,
	})

	service := &Service{
		ctx: &common.ServiceContext{Ctx: ctx, Config: cfg, DB: db, Redis: redis},
	}

	return service, nil
}

func (s *Service) Start() {
	startup.NewTaskStartup(s.ctx).Start()
	crontask.NewTask().Start()
	bounty.NewTaskBountyFactory(s.ctx).Start()
	bounty.NewTaskBounty(s.ctx).Start()
}
