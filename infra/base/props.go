package base

import (
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
	"resk/infra"
)

var props kvs.ConfigSource

//Props 对外暴露配置实例
func Props() kvs.ConfigSource {
	return props
}

//PropsStarter BaseStarter的实现
type PropsStarter struct {
	infra.BaseStarter
}

func (p *PropsStarter) Init(ctx infra.StarterContext) {
	props = ini.NewIniFileConfigSource("config.ini")
}
func (p *PropsStarter) Setup(context infra.StarterContext) {
	//TODO implement me
	panic("implement me")
}

func (p *PropsStarter) Start(context infra.StarterContext) {
	//TODO implement me
	panic("implement me")
}

func (p *PropsStarter) StartBlocking() bool {
	//TODO implement me
	panic("implement me")
}

func (p *PropsStarter) Stop(context infra.StarterContext) {
	//TODO implement me
	panic("implement me")
}

func (p *PropsStarter) PriorityGroup() infra.PriorityGroup {
	//TODO implement me
	panic("implement me")
}

func (p *PropsStarter) Priority() int {
	//TODO implement me
	panic("implement me")
}
