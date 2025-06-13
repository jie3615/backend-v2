package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	// RPCClientConf zrpc.RpcClientConf
	DB struct {
		DataSource string
	}
	JWT struct {
		Secret  string
		Expired int64
	}
	Minio struct {
		Bucket string
	}
	Redis struct {
		Host     string
		Password string
		DB       int
	}
}
