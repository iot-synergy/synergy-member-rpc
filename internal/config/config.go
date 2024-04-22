package config

import (
	"github.com/iot-synergy/synergy-common/config"

	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DatabaseConf config.DatabaseConf
	RedisConf    config.RedisConf
	CoreRpc      zrpc.RpcClientConf
	FcmRpc       zrpc.RpcClientConf
	MonDb        MonConfig
	AddxRpc      zrpc.RpcClientConf
}

type MonConfig struct {
	Url    string
	DbName string
}
