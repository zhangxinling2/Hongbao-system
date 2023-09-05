package base

import (
	"fmt"
	"github.com/tietang/props/kvs"
	"resk/infra"
)

var props kvs.ConfigSource

// Props 对外暴露配置实例
func Props() kvs.ConfigSource {
	return props
}

// PropsStarter BaseStarter的实现
type PropsStarter struct {
	infra.BaseStarter
}

func (p *PropsStarter) Init(ctx infra.StarterContext) {
	props = ctx.Props()
	fmt.Println("配置初始化")
}
