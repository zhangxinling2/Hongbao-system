package base

import (
	log "github.com/sirupsen/logrus"
	"net"
	"net/rpc"
	"reflect"
	"resk/infra"
)

var rpcServer *rpc.Server

func RpcServer() *rpc.Server {
	return rpcServer
}
func RpcRegister(ri interface{}) {
	typ := reflect.TypeOf(ri)
	log.Infof("goRPC Register: %s", typ.String())
	RpcServer().Register(ri)
}

type GoRPCStarter struct {
	infra.BaseStarter
}

func (g *GoRPCStarter) Init(ctx infra.StarterContext) {
	rpcServer = rpc.NewServer()
}
func (g *GoRPCStarter) Start(ctx infra.StarterContext) {
	port := Props().GetDefault("app.rpc.port", "8082")
	listner, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic("服务器监听失败")
	}
	log.Info("rpc服务开始监听" + port)
	go listner.Accept()
}
