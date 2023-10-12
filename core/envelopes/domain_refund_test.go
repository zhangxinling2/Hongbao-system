package envelopes

import (
	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	services "resk/services"
	"testing"
)

func TestExpiredEnvelopeDomain_Expire(t *testing.T) {
	domain := ExpiredEnvelopeDomain{}
	gDomain := goodsDomain{}
	iDomain := itemDomain{}
	domain.offset = 1
	Convey("红包退款", t, func() {
		Convey("查过期红包", func() {
			ok := domain.Next()
			So(ok, ShouldBeTrue)
			So(len(domain.expiredGoods), ShouldBeGreaterThan, 1)
			log.Info(domain.expiredGoods)
		})
		Convey("退款一个过期红包", func() {
			err := domain.ExpireOne(domain.expiredGoods[0])
			So(err, ShouldBeNil)
			good := gDomain.Get(domain.expiredGoods[0].EnvelopeNo)
			So(good.Status, ShouldEqual, services.OrderExpiredRefundSuccessful)
			items := iDomain.FindItems(domain.expiredGoods[0].EnvelopeNo)
			log.Info(items)
		})

	})
}
