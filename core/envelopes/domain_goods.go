package envelopes

import (
	"context"
	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	"github.com/tietang/dbx"
	"github.com/zhangxinling2/infra/base"
	services "github.com/zhangxinling2/resk/services"
	"time"
)

type goodsDomain struct {
	RedEnvelopeGoods
	item itemDomain
}

func (g *goodsDomain) createNo() {
	g.RedEnvelopeGoods.EnvelopeNo = ksuid.New().Next().String()
}
func (g *goodsDomain) Create(goods services.RedEnvelopeGoodsDTO) {
	g.RedEnvelopeGoods.FromDTO(goods)
	g.RedEnvelopeGoods.Blessing.Valid = true
	g.RedEnvelopeGoods.Username.Valid = true
	if g.RedEnvelopeGoods.EnvelopeType == services.GeneralEnvelopeType {
		g.RedEnvelopeGoods.Amount = goods.AmountOne.Mul(decimal.NewFromInt32(int32(goods.Quantity)))
	}
	if g.RedEnvelopeGoods.EnvelopeType == services.LuckyEnvelopeType {
		g.RedEnvelopeGoods.AmountOne = decimal.NewFromInt32(0)
	}
	g.RedEnvelopeGoods.RemainAmount = g.RedEnvelopeGoods.Amount
	g.RedEnvelopeGoods.RemainQuantity = g.RedEnvelopeGoods.Quantity
	g.RedEnvelopeGoods.ExpiredAt = time.Now().Add(24 * time.Hour)
	g.Status = services.OrderStatusCreated
	g.OrderType = services.OrderSending
	g.PayStatus = services.PayPaying
	g.createNo()
}
func (g *goodsDomain) Save(ctx context.Context) (id int64, err error) {
	err = base.ExecuteContext(ctx, func(runner *dbx.TxRunner) error {
		dao := EnvelopeDao{runner: runner}
		id, err = dao.Insert(&g.RedEnvelopeGoods)
		return err
	})
	return id, err
}
func (g *goodsDomain) CreateAndSave(ctx context.Context, goods services.RedEnvelopeGoodsDTO) (id int64, err error) {
	g.Create(goods)
	return g.Save(ctx)
}
func (g *goodsDomain) Get(envelopeNo string) (goods *RedEnvelopeGoods) {
	err := base.Tx(func(runner *dbx.TxRunner) error {
		dao := EnvelopeDao{runner: runner}
		goods = dao.GetOne(envelopeNo)
		return nil
	})
	if err != nil {
		return nil
	}
	return goods
}
