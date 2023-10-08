package service

import (
	"github.com/shopspring/decimal"
	"resk/infra/base"
	"time"
)

var redEnvelopeService RedEnvelopeService

func GetRedEnvelopeService() RedEnvelopeService {
	base.Check(redEnvelopeService)
	return redEnvelopeService
}

type RedEnvelopeService interface {
	//发红包
	SendOut(dto RedEnvelopeSendingDTO) (activity *RedEnvelopeActivity, err error)
	//收红包
	Receive(dto RedEnvelopeReceiveDTO) (item *RedEnvelopeItemDTO, err error)
	//退款
	Refund(envelopeNo string) (order *RedEnvelopeGoodsDTO)
	//查询红包订单
	Get(envelopeNo string) (order *RedEnvelopeGoodsDTO)
}

type RedEnvelopeSendingDTO struct {
	UserId   string          `json:"userId" validate:"require"`
	UserName string          `json:"userName" validate:"require"`
	Amount   decimal.Decimal `json:"amount" validate:"require,numeric"`
	Quantity int             `json:"quantity" validate:"require,numeric"`
	Blessing string          `json:"blessing"`
}
type RedEnvelopeReceiveDTO struct {
	EnvelopeNo   string `json:"envelopeNo" validate:"required"`   //红包编号,红包唯一标识
	RecvUsername string `json:"recvUsername" validate:"required"` //红包接收者用户名称
	RecvUserId   string `json:"recvUserId" validate:"required"`   //红包接收者用户编号
	AccountNo    string `json:"accountNo"`
}
type RedEnvelopeItemDTO struct {
	ItemNo       string          `json:"itemNo"`       //红包订单详情编号
	EnvelopeNo   string          `json:"envelopeNo"`   //订单编号 红包编号,红包唯一标识
	RecvUsername string          `json:"recvUsername"` //红包接收者用户名称
	RecvUserId   string          `json:"recvUserId"`   //红包接收者用户编号
	Amount       decimal.Decimal `json:"amount"`       //收到金额
	Quantity     int             `json:"quantity"`     //收到数量：对于收红包来说是1
	RemainAmount decimal.Decimal `json:"remainAmount"` //收到后红包剩余金额
	AccountNo    string          `json:"accountNo"`    //红包接收者账户ID
	PayStatus    int             `json:"payStatus"`    //支付状态：未支付，支付中，已支付，支付失败
	CreatedAt    time.Time       `json:"createdAt"`    //创建时间
	UpdatedAt    time.Time       `json:"updatedAt"`    //更新时间
	IsLuckiest   bool            `json:"isLuckiest"`
	Desc         string          `json:"desc"`
}
type RedEnvelopeGoodsDTO struct {
	EnvelopeNo       string          `json:"envelopeNo" validate:"require"`
	EnvelopeType     EnvelopeType    `json:"envelopeType" validate:"require"`
	UserName         string          `json:"userName" validate:"require"`
	UserId           string          `json:"userId" validate:"require"`
	Blessing         string          `json:"blessing"`
	Amount           decimal.Decimal `json:"amount" validate:"required,numeric"`
	AmountOne        decimal.Decimal `json:"amountOne"`
	Quantity         int             `json:"quantity" validate:"required,numeric"`
	RemainAmount     decimal.Decimal `json:"remainAmount" validate:"required,numeric"`
	RemainQuantity   int             `json:"remainQuantity" validate:"required,numeric"`
	ExpireAt         time.Time       `json:"expireAt"`
	Status           OrderStatus     `json:"status"`
	OrderType        OrderType       `json:"orderType"`
	PayStatus        PayStatus       `json:"payStatus"`
	CreateAt         time.Time       `json:"createAt"`
	UpdateAt         time.Time       `json:"updateAt"`
	AccountNo        string          `json:"accountNo"`
	OriginEnvelopeNo string          `json:"originEnvelopeNo"`
}
type RedEnvelopeActivity struct {
	RedEnvelopeGoodsDTO
	Link string `json:"link"`
}
