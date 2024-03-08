package svc

import (
	"github.com/iot-synergy/synergy-fcm/fcm"
	"github.com/iot-synergy/synergy-member-rpc/ent"
	"github.com/iot-synergy/synergy-member-rpc/internal/config"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	DB     *ent.Client
	Redis  redis.UniversalClient
	Fcm    fcm.Fcm
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := ent.NewClient(
		ent.Log(logx.Info), // logger
		ent.Driver(c.DatabaseConf.NewNoCacheDriver()),
		ent.Debug(), // debug mode
	)

	return &ServiceContext{
		Config: c,
		DB:     db,
		Redis:  c.RedisConf.MustNewUniversalRedis(),
		Fcm:    fcm.NewFcm(zrpc.MustNewClient(c.FcmRpc)),
	}
}
