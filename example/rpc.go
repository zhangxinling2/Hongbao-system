package main

import (
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"net/rpc"
	services "resk/services"
)

func main() {
	log.Info("开始dial")
	c, err := rpc.Dial("tcp", ":8082")
	if err != nil {
		panic(err)
	}
	in := services.RedEnvelopeSendingDTO{
		UserId:       "1001",
		UserName:     "测试用户",
		Amount:       decimal.NewFromInt32(22),
		Quantity:     1,
		Blessing:     "..",
		EnvelopeType: 1,
	}
	var out *services.RedEnvelopeActivity
	err = c.Call("EnvelopeRpc.SendOut", in, out)
	log.Info(err)
	log.Info(out)
}
