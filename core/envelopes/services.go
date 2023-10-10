package envelopes

import (
	"context"
	"errors"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"resk/infra/base"
	services "resk/services"
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
	ac := services.GetAccountService()
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
	as := services.GetAccountService()
	account := as.GetEnvelopeAccountByUserId(dto.RecvUserId)
	if account != nil {
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
	//TODO implement me
	panic("implement me")
}
