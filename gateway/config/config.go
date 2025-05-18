package config

import (
	"gateway/config/sever"
	"gateway/config/theetcd"
	"gateway/config/thelog"
)

type Config struct {
	Gates []sever.GateWaySeverConf
	Log   thelog.LoggerConf
	etcd  theetcd.EtcdConf
}
