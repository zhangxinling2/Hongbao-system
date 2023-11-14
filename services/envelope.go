package service

import (
	"github.com/shopspring/decimal"
	"github.com/zhangxinling2/infra/base"
	"time"
)

var IredEnvelopeService RedEnvelopeService

func GetRedEnvelopeService() RedEnvelopeService {
	base.Check(IredEnvelopeService)
	return IredEnvelopeService
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
	//查询用户已经发送的红包列表
	ListSent(userId string, page, size int) (orders []*RedEnvelopeGoodsDTO)
	ListReceived(userId string, page, size int) (items []*RedEnvelopeItemDTO)
	//查询用户已经抢到的红包列表
	ListReceivable(page, size int) (orders []*RedEnvelopeGoodsDTO)
	ListItems(envelopeNo string) (items []*RedEnvelopeItemDTO)
}

type RedEnvelopeSendingDTO struct {
	UserId       string          `json:"userId" validate:"required"`
	UserName     string          `json:"userName" validate:"required"`
	Amount       decimal.Decimal `json:"amount" validate:"required,numeric"`
	Quantity     int             `json:"quantity" validate:"required,numeric"`
	Blessing     string          `json:"blessing"`
	EnvelopeType EnvelopeType    `json:"envelopeType"`
}

func (r *RedEnvelopeSendingDTO) ToGoods() *RedEnvelopeGoodsDTO {
	goods := &RedEnvelopeGoodsDTO{
		EnvelopeType: r.EnvelopeType,
		UserName:     r.UserName,
		UserId:       r.UserId,
		Blessing:     r.Blessing,
		Amount:       r.Amount,
		Quantity:     r.Quantity,
	}
	return goods
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
	EnvelopeNo       string          `json:"envelopeNo" validate:"required"`
	EnvelopeType     EnvelopeType    `json:"envelopeType" validate:"required"`
	UserName         string          `json:"userName" validate:"required"`
	UserId           string          `json:"userId" validate:"required"`
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

func (r *RedEnvelopeActivity) CopyTo(res *RedEnvelopeActivity) {
	res = r
}
