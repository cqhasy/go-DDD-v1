package etcdx

import (
	"gateway/public"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func InitEtcd(path []string) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   path,
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		panic(err)
	}
	public.EtcdClient = cli
}
