package envelopes

import (
	"context"
	"github.com/tietang/dbx"
	"github.com/zhangxinling2/account/core/accounts"
	acservices "github.com/zhangxinling2/account/services"
	"github.com/zhangxinling2/infra/base"
	services "github.com/zhangxinling2/resk/services"
	"path"
)

func (g *goodsDomain) SendOut(goods services.RedEnvelopeGoodsDTO) (activity *services.RedEnvelopeActivity, err error) {
	g.Create(goods)
	activity = new(services.RedEnvelopeActivity)
	link := base.GetEnvelopeLink()
	domain := base.GetEnvelopeDomain()
	activity.Link = path.Join(domain, link, g.EnvelopeNo)
	accountDomain := accounts.NewAccountDomain()
	err = base.Tx(func(runner *dbx.TxRunner) error {
		ctx := base.WithValueContext(context.Background(), runner)
		id, err := g.Save(ctx)
		if id <= 0 || err != nil {
			return err
		}
		body := acservices.TradeParticipator{
			AccountNo: goods.AccountNo,
			UserId:    goods.UserId,
			UserName:  goods.UserName,
		}
		systemAccount := base.GetSystemAccount()
		target := acservices.TradeParticipator{
			AccountNo: systemAccount.AccountNo,
			UserId:    systemAccount.UserId,
			UserName:  systemAccount.UserName,
		}
		transfer := acservices.AccountTransferDTO{
			TradeNo:     g.EnvelopeNo,
			TradeBody:   body,
			TradeTarget: target,
			AmountStr:   g.Amount.String(),
			Amount:      g.Amount,
			ChangeType:  acservices.EnvelopeOutgoing,
			ChangeFlag:  acservices.FlagTransferOut,
			Decs:        "红包金额支付",
		}
		status, err := accountDomain.TransferWithContextTx(ctx, transfer)
		if status != acservices.TransferSuccess {
			return err
		}
		transfer = acservices.AccountTransferDTO{
			TradeNo:     g.EnvelopeNo,
			TradeBody:   target,
			TradeTarget: body,
			AmountStr:   g.Amount.String(),
			Amount:      g.Amount,
			ChangeType:  acservices.EnvelopeIncoming,
			ChangeFlag:  acservices.FlagTransferIn,
			Decs:        "红包金额转入",
		}
		status, err = accountDomain.TransferWithContextTx(ctx, transfer)
		if status != acservices.TransferSuccess {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	activity.RedEnvelopeGoodsDTO = *g.RedEnvelopeGoods.ToDTO()
	return activity, err
}
