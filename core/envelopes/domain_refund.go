package envelopes

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
	"github.com/zhangxinling2/account/core/accounts"
	acservices "github.com/zhangxinling2/account/services"
	"github.com/zhangxinling2/infra/base"
	services "github.com/zhangxinling2/resk/services"
)

type ExpiredEnvelopeDomain struct {
	expiredGoods []RedEnvelopeGoods
	offset       int
}

var size int = 100

func (e *ExpiredEnvelopeDomain) Next() (ok bool) {
	base.Tx(func(runner *dbx.TxRunner) error {
		dao := EnvelopeDao{runner: runner}
		e.expiredGoods = dao.FindExpired(e.offset, size)
		if len(e.expiredGoods) > 0 {
			e.offset += size
			ok = true
		}
		return nil
	})
	return ok
}
func (e *ExpiredEnvelopeDomain) Expire() (err error) {
	for e.Next() {
		for _, goods := range e.expiredGoods {
			if goods.OrderType == services.OrderSending {
				err = e.ExpireOne(goods)
				if err != nil {
					log.Error(err)
				}
			}
		}
	}
	return err
}
func (e *ExpiredEnvelopeDomain) ExpireOne(goods RedEnvelopeGoods) (err error) {
	//创建退款订单
	refund := goods
	refund.OrderType = services.OrderRefund
	//refund.RemainAmount = goods.RemainAmount.Mul(decimal.NewFromFloat(-1))
	//refund.RemainQuantity = -goods.RemainQuantity
	refund.PayStatus = services.Refunding
	refund.Status = services.OrderStatusExpire
	refund.OriginEnvelopeNo = goods.EnvelopeNo
	domain := goodsDomain{RedEnvelopeGoods: refund}
	domain.createNo()
	err = base.Tx(func(runner *dbx.TxRunner) error {
		txCtx := base.WithValueContext(context.Background(), runner)
		id, err := domain.Save(txCtx)
		if id <= 0 || err != nil {
			return errors.New("创建退款订单失败")
		}
		//更改原订单状态，与创建退款订单同一事务
		dao := EnvelopeDao{runner: runner}
		rows, err := dao.UpdateOrderStatus(goods.EnvelopeNo, services.OrderStatusExpire)
		if err != nil || rows <= 0 {
			return errors.New("更改原订单状态失败")
		}
		return nil
	})
	//调用资金接口进行转账
	acDomain := accounts.NewAccountDomain()
	sc := base.GetSystemAccount()
	ac := acDomain.GetEnvelopeAccountByUserId(goods.UserId)
	body := acservices.TradeParticipator{
		AccountNo: sc.AccountNo,
		UserId:    sc.UserId,
		UserName:  sc.UserName,
	}
	target := acservices.TradeParticipator{
		AccountNo: ac.AccountNo,
		UserId:    ac.UserId,
		UserName:  ac.UserName,
	}
	transfer := acservices.AccountTransferDTO{
		TradeNo:     domain.RedEnvelopeGoods.EnvelopeNo,
		TradeBody:   body,
		TradeTarget: target,
		AmountStr:   "",
		Amount:      goods.RemainAmount,
		ChangeType:  acservices.EnvelopeExpire,
		ChangeFlag:  acservices.FlagTransferOut,
		Decs:        "红包过期退款",
	}
	status, err := acservices.GetAccountService().Transfer(transfer)
	if status != acservices.TransferSuccess || err != nil {
		return err
	}
	//更新原订单状态
	err = base.Tx(func(runner *dbx.TxRunner) error {
		dao := EnvelopeDao{runner: runner}
		rows, err := dao.UpdateOrderStatus(goods.EnvelopeNo, services.OrderExpiredRefundSuccessful)
		if err != nil || rows <= 0 {
			return errors.New("更改原订单状态失败" + err.Error())
		}
		rows, err = dao.UpdateOrderStatus(domain.RedEnvelopeGoods.EnvelopeNo, services.OrderExpiredRefundSuccessful)
		if err != nil || rows <= 0 {
			return errors.New("更改退款订单状态失败" + err.Error())
		}
		return nil
	})
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
