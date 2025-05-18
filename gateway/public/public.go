package public

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger
var EtcdClient *clientv3.Client
