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
	err = sendout(c)
}
func sendout(c *rpc.Client) error {
	in := services.RedEnvelopeSendingDTO{
		UserId:       "1001",
		UserName:     "测试用户",
		Amount:       decimal.NewFromInt32(22),
		Quantity:     1,
		Blessing:     "..",
		EnvelopeType: 1,
	}
	var out *services.RedEnvelopeActivity
	err := c.Call("EnvelopeRpc.SendOut", in, out)
	log.Info(err)
	log.Info(out)
	return err
}
func receive(c *rpc.Client) error {
	in := services.RedEnvelopeReceiveDTO{
		EnvelopeNo:   "",
		RecvUserId:   "",
		RecvUsername: "",
		AccountNo:    "",
	}
	out := &services.RedEnvelopeItemDTO{}
	err := c.Call("Envelope.Receive", in, out)
	if err != nil {
		log.Panic(err)
	}
	log.Info(out)
	return err
}
