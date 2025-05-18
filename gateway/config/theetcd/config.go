package theetcd

import "gateway/tools/etcdx"

type EtcdConf struct {
	path []string
}

func NewEtcdConf(path []string) *EtcdConf {
	return &EtcdConf{path}
}
func (e *EtcdConf) Path() []string {
	return e.path
}
func (e *EtcdConf) Init() {
	etcdx.InitEtcd(e.path)
}
