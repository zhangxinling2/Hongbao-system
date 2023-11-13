package gorpc

import (
	"github.com/zhangxinling2/infra"
	"github.com/zhangxinling2/infra/base"
)

type GoRpcApiStarter struct {
	infra.BaseStarter
}

func (g *GoRpcApiStarter) Init(ctx infra.StarterContext) {
	base.RpcRegister(new(EnvelopeRPC))
}
