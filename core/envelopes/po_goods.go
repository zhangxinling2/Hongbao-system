package envelopes

import (
	"database/sql"
	"github.com/shopspring/decimal"
	services "resk/services"
	"time"
)

type RedEnvelopeGoods struct {
	Id             int64                 `db:"id,omitempty"`         //自增ID
	EnvelopeNo     string                `db:"envelope_no,uni"`      //红包编号,红包唯一标识
	EnvelopeType   services.EnvelopeType `db:"envelope_type"`        //红包类型：普通红包，碰运气红包
	Username       sql.NullString        `db:"username"`             //用户名称
	UserId         string                `db:"user_id"`              //用户编号, 红包所属用户
	Blessing       sql.NullString        `db:"blessing"`             //祝福语
	Amount         decimal.Decimal       `db:"amount"`               //红包总金额
	AmountOne      decimal.Decimal       `db:"amount_one"`           //单个红包金额，碰运气红包无效
	Quantity       int                   `db:"quantity"`             //红包总数量
	RemainAmount   decimal.Decimal       `db:"remain_amount"`        //红包剩余金额额
	RemainQuantity int                   `db:"remain_quantity"`      //红包剩余数量
	ExpiredAt      time.Time             `db:"expired_at"`           //过期时间
	Status         services.OrderStatus  `db:"status"`               //红包状态：0红包初始化，1启用，2失效
	OrderType      services.OrderType    `db:"order_type"`           //订单类型：发布单、退款单
	PayStatus      services.PayStatus    `db:"pay_status"`           //支付状态：未支付，支付中，已支付，支付失败
	CreatedAt      time.Time             `db:"created_at,omitempty"` //创建时间
	UpdatedAt      time.Time             `db:"updated_at,omitempty"` //更新时间
}

func (r *RedEnvelopeGoods) FromDTO(dto services.RedEnvelopeGoodsDTO) {
	r.EnvelopeNo = dto.EnvelopeNo
	r.EnvelopeType = dto.EnvelopeType
	r.Username.String = dto.UserName
	r.UserId = dto.UserId
	r.Blessing.String = dto.Blessing
	r.Amount = dto.Amount
	r.Quantity = dto.Quantity
	r.OrderType = dto.OrderType
	r.Status = dto.Status
	r.PayStatus = dto.PayStatus
}
