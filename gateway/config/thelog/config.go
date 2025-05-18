package thelog

import "gateway/tools/logger"

type LoggerConf struct {
	path string
}

func NewLoggerConf(path string) *LoggerConf {
	return &LoggerConf{path: path}
}
func (conf *LoggerConf) GetPath() string {
	return conf.path
}
func (conf *LoggerConf) Init() {
	logger.Initlog(conf.path)
}
