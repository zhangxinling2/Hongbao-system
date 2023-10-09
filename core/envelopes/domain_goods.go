package envelopes

import (
	"context"
	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	"github.com/tietang/dbx"
	"resk/infra/base"
	services "resk/services"
	"time"
)

type goodsDomain struct {
	RedEnvelopeGoods
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
