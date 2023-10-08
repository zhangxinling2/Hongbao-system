package envelopes

import (
	"database/sql"
	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tietang/dbx"
	"resk/infra/base"
	_ "resk/testx"
	"testing"
	"time"
)

func TestEnvelopeDao_GetOne(t *testing.T) {
	err := base.Tx(func(runner *dbx.TxRunner) error {
		dao := EnvelopeDao{runner: runner}
		envelopeAmount := decimal.NewFromFloat(100)
		envelopeAmountOne := decimal.NewFromFloat(50)
		h, _ := time.ParseDuration("24h")
		expire := time.Now().Add(h)
		a := &RedEnvelopeGoods{
			EnvelopeNo:     "2",
			EnvelopeType:   1,
			Username:       sql.NullString{Valid: true, String: "红包测试用户"},
			UserId:         ksuid.New().Next().String(),
			Amount:         envelopeAmount,
			AmountOne:      envelopeAmountOne,
			Quantity:       2,
			RemainAmount:   envelopeAmount,
			RemainQuantity: 2,
			ExpiredAt:      expire,
			Status:         0,
		}

		_, err := dao.Insert(a)
		rs := dao.GetOne(a.EnvelopeNo)

		Convey("更新余额", t, func() {
			i, err := dao.UpdateBalance(a.EnvelopeNo, envelopeAmountOne)
			So(err, ShouldBeNil)
			So(i, ShouldEqual, 1)
			rs = dao.GetOne(a.EnvelopeNo)
			So(rs.RemainAmount.String(), ShouldEqual, envelopeAmountOne.String())
			So(rs.RemainQuantity, ShouldEqual, 1)
		})
		Convey("更新状态", t, func() {
			i, err := dao.UpdateOrderStatus(a.EnvelopeNo, 2)
			So(err, ShouldBeNil)
			So(i, ShouldEqual, 1)
			rs = dao.GetOne(a.EnvelopeNo)
			So(rs.Status, ShouldEqual, 2)

		})

		return err
	})
	logrus.Error(err)
}
