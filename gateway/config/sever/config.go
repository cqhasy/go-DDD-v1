package sever

type GateWaySeverConf struct {
	Addr   string
	Path   string
	Target string
	Method string
}

func (conf *GateWaySeverConf) GetUrl() string {
	return conf.Addr + "/" + conf.Path
}
func (conf *GateWaySeverConf) GetTarget() string {
	return conf.Target
}
func (conf *GateWaySeverConf) GetMethod() string {
	return conf.Method
}
