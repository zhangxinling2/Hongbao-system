package jobs

import (
	"resk/infra"
	"resk/infra/base"
	services "resk/services"
	"time"
)

type RefundExpiredJobStarter struct {
	infra.BaseStarter
	ExpiredGoods []services.RedEnvelopeGoodsDTO
	ticker       *time.Ticker
}

func (r *RefundExpiredJobStarter) Init(ctx infra.StarterContext) {
	dur := base.Props().GetDurationDefault("jobs.refund.internal", time.Minute)
	r.ticker = time.NewTicker(dur)
}
func (r *RefundExpiredJobStarter) Start(ctx infra.StarterContext) {
	go func() {
		for {
			<-r.ticker.C
		}
	}()
}
func (r *RefundExpiredJobStarter) Stop(ctx infra.StarterContext) {
	r.ticker.Stop()
}
