package envelopes

import (
	services "resk/services"
	"sync"
)

var _ services.RedEnvelopeService = new(envelopeService)
var once sync.Once

func init() {
	once.Do(func() {
		services.IredEnvelopeService = new(envelopeService)
	})
}

type envelopeService struct {
}

func (e *envelopeService) SendOut(dto services.RedEnvelopeSendingDTO) (activity *services.RedEnvelopeActivity, err error) {
	domain := new(goodsDomain)

}

func (e *envelopeService) Receive(dto services.RedEnvelopeReceiveDTO) (item *services.RedEnvelopeItemDTO, err error) {
	//TODO implement me
	panic("implement me")
}

func (e *envelopeService) Refund(envelopeNo string) (order *services.RedEnvelopeGoodsDTO) {
	//TODO implement me
	panic("implement me")
}

func (e *envelopeService) Get(envelopeNo string) (order *services.RedEnvelopeGoodsDTO) {
	//TODO implement me
	panic("implement me")
}
