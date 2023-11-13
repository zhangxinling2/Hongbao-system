package envelopes

import (
	"context"
	"database/sql"
	"errors"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
	"resk/core/accounts"
	"github.com/zhangxinling2/infra/algo"
	"github.com/zhangxinling2/infra/base"
	services "resk/services"
)

var mul = decimal.NewFromFloat(100.0)

func (d *goodsDomain) Receive(ctx context.Context, dto services.RedEnvelopeReceiveDTO) (item *services.RedEnvelopeItemDTO, err error) {
	d.preCreateItem(dto)
	good := d.Get(dto.EnvelopeNo)
	if good.RemainQuantity <= 0 || good.RemainAmount.Cmp(decimal.NewFromFloat(0)) <= 0 {
		log.Errorf("没有足够的红包和金额了: %+v", good)
		return nil, errors.New("没有足够的红包和金额了")
	}
	amount, err := d.nextAmount(good)
	log.Info("amount" + amount.String())
	if err != nil {
		log.Error(err)
		return nil, err
	}
	err = base.Tx(func(runner *dbx.TxRunner) error {
		dao := EnvelopeDao{runner: runner}
		i, err := dao.UpdateBalance(good.EnvelopeNo, amount)
		if err != nil || i <= 0 {
			return errors.New("没有足够的红包和金额了")
		}
		d.item.Quantity = 1
		d.item.Amount = amount
		d.item.PayStatus = int(services.PayPaying)
		d.item.AccountNo = dto.AccountNo
		d.item.RemainAmount = good.RemainAmount.Sub(amount)
		txCtx := base.WithValueContext(ctx, runner)
		_, err = d.item.Save(txCtx)
		if err != nil {
			log.Error(err)
			return err
		}
		status, err := d.transfer(ctx, dto)
		if status == services.TransferSuccess {
			return nil
		} else {
			return err
		}
		return err
	})
	if err != nil {
		return nil, err
	}
	return d.item.ToDTO(), err
}
func (d *goodsDomain) preCreateItem(dto services.RedEnvelopeReceiveDTO) {
	d.item.EnvelopeNo = dto.EnvelopeNo
	d.item.RecvUsername = sql.NullString{
		String: dto.RecvUsername,
		Valid:  true,
	}
	d.item.RecvUserId = dto.RecvUserId
	d.item.AccountNo = dto.AccountNo
	d.item.createItemNo()
}
func (d *goodsDomain) nextAmount(goods *RedEnvelopeGoods) (amount decimal.Decimal, err error) {
	if goods.RemainQuantity == 1 {
		return goods.RemainAmount, nil
	}
	if goods.EnvelopeType == services.GeneralEnvelopeType {
		return goods.AmountOne, nil
	} else if goods.EnvelopeType == services.LuckyEnvelopeType {
		cent := goods.RemainAmount.Mul(mul).IntPart()
		next := algo.DoubleAverage(int64(goods.RemainQuantity), cent)
		amount = decimal.NewFromFloat(float64(next)).Div(mul)
	} else {
		return decimal.NewFromInt(0), errors.New("不支持的红包类型")
	}
	return amount, nil
}
func (d *goodsDomain) transfer(ctx context.Context, dto services.RedEnvelopeReceiveDTO) (status services.TransferStatus, err error) {
	sa := base.GetSystemAccount()
	body := services.TradeParticipator{
		AccountNo: sa.AccountNo,
		UserId:    sa.UserId,
		UserName:  sa.UserName,
	}
	target := services.TradeParticipator{
		AccountNo: dto.AccountNo,
		UserId:    dto.RecvUserId,
		UserName:  dto.RecvUsername,
	}
	transfer := services.AccountTransferDTO{
		TradeNo:     dto.EnvelopeNo,
		TradeBody:   body,
		TradeTarget: target,
		AmountStr:   "",
		Amount:      d.item.Amount,
		ChangeType:  services.EnvelopeOutgoing,
		ChangeFlag:  services.FlagTransferOut,
		Decs:        "红包扣减" + dto.EnvelopeNo,
	}
	domain := accounts.NewAccountDomain()
	status, err = domain.TransferWithContextTx(ctx, transfer)
	if err != nil || status != services.TransferSuccess {
		return status, err
	}
	transfer = services.AccountTransferDTO{
		TradeNo:     dto.EnvelopeNo,
		TradeBody:   target,
		TradeTarget: body,
		AmountStr:   "",
		Amount:      d.item.Amount,
		ChangeType:  services.EnvelopeIncoming,
		ChangeFlag:  services.FlagTransferIn,
		Decs:        "收红包" + dto.EnvelopeNo,
	}
	return domain.TransferWithContextTx(ctx, transfer)
}
