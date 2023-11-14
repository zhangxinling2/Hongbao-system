package envelopes

import (
	"context"
	"errors"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	acservices "github.com/zhangxinling2/account/services"
	"github.com/zhangxinling2/infra/base"
	services "github.com/zhangxinling2/resk/services"
	"sync"
)

var _ services.RedEnvelopeService = new(redEnvelopeService)
var once sync.Once

func init() {
	once.Do(func() {
		services.IredEnvelopeService = new(redEnvelopeService)
	})
}

type redEnvelopeService struct {
}

func (e *redEnvelopeService) SendOut(dto services.RedEnvelopeSendingDTO) (activity *services.RedEnvelopeActivity, err error) {
	err = base.ValidateStruct(&dto)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	ac := acservices.GetAccountService()
	account := ac.GetEnvelopeAccountByUserId(dto.UserId)
	if account == nil {
		return nil, errors.New("账户不存在")
	}
	goods := dto.ToGoods()
	goods.AccountNo = account.AccountNo
	if goods.Blessing == "" {
		goods.Blessing = services.DefaultBlessing
	}
	if goods.EnvelopeType == services.GeneralEnvelopeType {
		goods.AmountOne = goods.Amount
		goods.Amount = decimal.Decimal{}
	}
	domain := new(goodsDomain)
	activity, err = domain.SendOut(*goods)
	if err != nil {
		log.Error(err)
	}

	return activity, err
}

func (e *redEnvelopeService) Receive(dto services.RedEnvelopeReceiveDTO) (item *services.RedEnvelopeItemDTO, err error) {
	err = base.ValidateStruct(&dto)
	as := acservices.GetAccountService()
	account := as.GetEnvelopeAccountByUserId(dto.RecvUserId)
	if account == nil {
		return nil, errors.New("红包账户不存在:user_id=" + dto.RecvUserId)
	}
	domain := new(goodsDomain)
	return domain.Receive(context.Background(), dto)
}

func (e *redEnvelopeService) Refund(envelopeNo string) (order *services.RedEnvelopeGoodsDTO) {
	//TODO implement me
	panic("implement me")
}

func (e *redEnvelopeService) Get(envelopeNo string) (order *services.RedEnvelopeGoodsDTO) {
	domain := goodsDomain{}
	po := domain.Get(envelopeNo)
	if po == nil {
		return order
	}
	return po.ToDTO()
}
func (r *redEnvelopeService) ListSent(userId string, page, size int) (orders []*services.RedEnvelopeGoodsDTO) {
	domain := new(goodsDomain)
	pos := domain.FindByUser(userId, page, size)
	orders = make([]*services.RedEnvelopeGoodsDTO, 0, len(pos))
	for _, p := range pos {
		orders = append(orders, p.ToDTO())
	}

	return
}

func (r *redEnvelopeService) ListReceivable(page, size int) (orders []*services.RedEnvelopeGoodsDTO) {
	domain := new(goodsDomain)

	pos := domain.ListReceivable(page, size)
	orders = make([]*services.RedEnvelopeGoodsDTO, 0, len(pos))
	for _, p := range pos {
		if p.RemainQuantity > 0 {
			orders = append(orders, p.ToDTO())
		}
	}
	return
}

func (r *redEnvelopeService) ListReceived(userId string, page, size int) (items []*services.RedEnvelopeItemDTO) {
	domain := new(goodsDomain)
	pos := domain.ListReceived(userId, page, size)
	items = make([]*services.RedEnvelopeItemDTO, 0, len(pos))
	if len(pos) == 0 {
		return items
	}
	for _, p := range pos {
		items = append(items, p.ToDTO())
	}
	return
}

func (r *redEnvelopeService) ListItems(envelopeNo string) (items []*services.RedEnvelopeItemDTO) {
	domain := itemDomain{}
	return domain.FindItems(envelopeNo)
}
