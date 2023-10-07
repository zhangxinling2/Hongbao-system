package service

import "github.com/shopspring/decimal"

type RedEnvelopeService interface {
	//发红包
	SendOut()
	//收红包
	Receive()
	//退款
	Refund()
	//查询红包订单
	Get()
}

type RedEnvelopeSendingDTO struct {
	UserId   string          `json:"userId" validate:"require"`
	UserName string          `json:"userName" validate:"require"`
	Amount   decimal.Decimal `json:"amount" validate:"require,numeric"`
	Quantity int             `json:"quantity" validate:"require,numeric"`
}
