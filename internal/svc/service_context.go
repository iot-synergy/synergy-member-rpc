package svc

import (
	"github.com/iot-synergy/synergy-activation-code-rpc/activationcoderpcclient"
	"github.com/iot-synergy/synergy-addx-proxy/synergy_addx_proxy_client"
	"github.com/iot-synergy/synergy-fcm/fcm"
	"github.com/iot-synergy/synergy-member-rpc/ent"
	"github.com/iot-synergy/synergy-member-rpc/internal/config"
	model "github.com/iot-synergy/synergy-member-rpc/storage/alarm_config"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config                   config.Config
	DB                       *ent.Client
	Redis                    redis.UniversalClient
	Fcm                      fcm.Fcm
	AlarmConfigModel         model.Alarm_configModel
	AddxRpc                  synergy_addx_proxy_client.SynergyAddxProxy
	SynergyActivationCodeRpc activationcoderpcclient.Activationcoderpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := ent.NewClient(
		ent.Log(logx.Info), // logger
		ent.Driver(c.DatabaseConf.NewNoCacheDriver()),
		ent.Debug(), // debug mode
	)

	return &ServiceContext{
		Config:                   c,
		DB:                       db,
		Redis:                    c.RedisConf.MustNewUniversalRedis(),
		Fcm:                      fcm.NewFcm(zrpc.MustNewClient(c.FcmRpc)),
		AlarmConfigModel:         model.NewAlarm_configModel(c.MonDb.Url, c.MonDb.DbName, "alarm_config"),
		AddxRpc:                  synergy_addx_proxy_client.NewSynergyAddxProxy(zrpc.MustNewClient(c.AddxRpc)),
		SynergyActivationCodeRpc: activationcoderpcclient.NewActivationcoderpc(zrpc.MustNewClient(c.SynergyActivationCodeRpc)),
	}
}
