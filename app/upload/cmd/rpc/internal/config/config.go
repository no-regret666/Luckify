package config

import (
	"github.com/SpectatorNan/gorm-zero/gormc/config/mysql"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	MySQL   mysql.Mysql
	Cache   cache.CacheConf
	OssConf struct {
		Endpoint  string // OSS服务地址
		AccessKey string
		SecretKey string
		Bucket    string
		Region    string
		UseSSL    bool
	}
}
