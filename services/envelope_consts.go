package service

const (
	DefaultBlessing string = "恭喜发财"
)

// 订单类型
type OrderType int

const (
	OrderSending OrderType = 1
	OrderRefund  OrderType = 2
)

// 支付状态
type PayStatus int

const (
	PayNothing PayStatus = 1
	PayPaying  PayStatus = 2
	PayPayed   PayStatus = 3
	PayFailure PayStatus = 4

	RefundNothing PayStatus = 61
	Refunding     PayStatus = 62
	Refunded      PayStatus = 63
	RefundFailure PayStatus = 64
)

type OrderStatus int

const (
	OrderStatusCreated           OrderStatus = 1
	OrderStatusSending           OrderStatus = 2
	OrderStatusExpire            OrderStatus = 3
	OrderStatusDisabled          OrderStatus = 4
	OrderExpiredRefundSuccessful OrderStatus = 5
	OrderExpiredRefundFailure    OrderStatus = 6
)

// 红包类型：普通红包，碰运气红包
type EnvelopeType int

const (
	GeneralEnvelopeType = 1
	LuckyEnvelopeType   = 2
)
