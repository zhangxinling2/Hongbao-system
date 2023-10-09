package envelopes

import (
	"context"
	"github.com/tietang/dbx"
	"path"
	"resk/core/accounts"
	"resk/infra/base"
	services "resk/services"
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
		body := services.TradeParticipator{
			AccountNo: goods.AccountNo,
			UserId:    goods.UserId,
			UserName:  goods.UserName,
		}
		systemAccount := base.GetSystemAccount()
		target := services.TradeParticipator{
			AccountNo: systemAccount.AccountNo,
			UserId:    systemAccount.UserId,
			UserName:  systemAccount.UserName,
		}
		transfer := services.AccountTransferDTO{
			TradeNo:     g.EnvelopeNo,
			TradeBody:   body,
			TradeTarget: target,
			AmountStr:   g.Amount.String(),
			Amount:      g.Amount,
			ChangeType:  services.EnvelopeOutgoing,
			ChangeFlag:  services.FlagTransferOut,
			Decs:        "红包金额支付",
		}
		status, err := accountDomain.TransferWithContextTx(ctx, transfer)
		if status != services.TransferSuccess {
			return err
		}
		transfer = services.AccountTransferDTO{
			TradeNo:     g.EnvelopeNo,
			TradeBody:   target,
			TradeTarget: body,
			AmountStr:   g.Amount.String(),
			Amount:      g.Amount,
			ChangeType:  services.EnvelopeIncoming,
			ChangeFlag:  services.FlagTransferIn,
			Decs:        "红包金额转入",
		}
		status, err = accountDomain.TransferWithContextTx(ctx, transfer)
		if status != services.TransferSuccess {
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
