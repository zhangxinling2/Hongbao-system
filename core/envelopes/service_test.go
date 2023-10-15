package envelopes

import (
	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	. "github.com/smartystreets/goconvey/convey"
	service "resk/services"
	_ "resk/testx"
	"testing"
)

func TestRedEnvelopeService_SendOut(t *testing.T) {
	re := service.GetRedEnvelopeService()
	account := service.AccountCreatedDTO{
		UserId:      ksuid.New().Next().String(),
		UserName:    "测试账户",
		AccountName: "测试账户",
		AccountType: int(service.EnvelopeAccountType),
		CurrentCode: "CNY",
		Amount:      "1000",
	}
	ac := service.GetAccountService()
	Convey("创建账户", t, func() {
		acDTO, err := ac.CreateAccount(account)
		So(err, ShouldBeNil)
		So(acDTO, ShouldNotBeNil)
	})
	Convey("发送红包", t, func() {
		Convey("发送普通红包", func() {
			goods := service.RedEnvelopeSendingDTO{
				UserId:       account.UserId,
				UserName:     account.UserName,
				EnvelopeType: service.GeneralEnvelopeType,
				Amount:       decimal.NewFromFloat(8.88),
				Quantity:     10,
				Blessing:     service.DefaultBlessing,
			}
			activity, err := re.SendOut(goods)
			So(err, ShouldBeNil)
			So(activity, ShouldNotBeNil)
			So(activity.Link, ShouldNotBeEmpty)
			So(activity.RedEnvelopeGoodsDTO, ShouldNotBeNil)
			dto := activity.RedEnvelopeGoodsDTO
			So(dto.UserName, ShouldEqual, goods.UserName)
			So(dto.UserId, ShouldEqual, goods.UserId)
			So(dto.Quantity, ShouldEqual, goods.Quantity)
			So(dto.Amount.String(), ShouldEqual, goods.Amount.Mul(decimal.NewFromInt(int64(goods.Quantity))).String())
		})
		Convey("发碰运气红包", func() {
			goods := service.RedEnvelopeSendingDTO{
				UserId:       account.UserId,
				UserName:     account.UserName,
				EnvelopeType: service.LuckyEnvelopeType,
				Amount:       decimal.NewFromFloat(88.8),
				Quantity:     10,
				Blessing:     service.DefaultBlessing,
			}
			activity, err := re.SendOut(goods)
			So(err, ShouldBeNil)
			So(activity, ShouldNotBeNil)
			So(activity.Link, ShouldNotBeEmpty)
			So(activity.RedEnvelopeGoodsDTO, ShouldNotBeNil)
			dto := activity.RedEnvelopeGoodsDTO
			So(dto.UserName, ShouldEqual, goods.UserName)
			So(dto.UserId, ShouldEqual, goods.UserId)
			So(dto.Quantity, ShouldEqual, goods.Quantity)
			So(dto.Amount, ShouldEqual, goods.Amount)
		})
	})
}
func TestRedEnvelopeService_SendOut_Failure(t *testing.T) {
	re := service.GetRedEnvelopeService()
	account := service.AccountCreatedDTO{
		UserId:      ksuid.New().Next().String(),
		UserName:    "测试账户F",
		AccountName: "测试账户F",
		AccountType: int(service.EnvelopeAccountType),
		CurrentCode: "CNY",
		Amount:      "30",
	}
	ac := service.GetAccountService()
	Convey("创建账户", t, func() {
		acDTO, err := ac.CreateAccount(account)
		So(err, ShouldBeNil)
		So(acDTO, ShouldNotBeNil)
	})
	Convey("发送红包失败", t, func() {
		Convey("发送普通红包失败", func() {
			goods := service.RedEnvelopeSendingDTO{
				UserId:       account.UserId,
				UserName:     account.UserName,
				EnvelopeType: service.GeneralEnvelopeType,
				Amount:       decimal.NewFromFloat(8.88),
				Quantity:     10,
				Blessing:     service.DefaultBlessing,
			}
			activity, err := re.SendOut(goods)
			So(err, ShouldNotBeNil)
			So(activity, ShouldBeNil)
			a := ac.GetEnvelopeAccountByUserId(account.UserId)
			So(a.Balance.String(), ShouldEqual, account.Amount)
		})
		Convey("发碰运气红包失败", func() {
			goods := service.RedEnvelopeSendingDTO{
				UserId:       account.UserId,
				UserName:     account.UserName,
				EnvelopeType: service.LuckyEnvelopeType,
				Amount:       decimal.NewFromFloat(88.8),
				Quantity:     10,
				Blessing:     service.DefaultBlessing,
			}
			activity, err := re.SendOut(goods)
			So(err, ShouldNotBeNil)
			So(activity, ShouldBeNil)
			a := ac.GetEnvelopeAccountByUserId(account.UserId)
			So(a.Balance.String(), ShouldEqual, account.Amount)
		})
	})
}
