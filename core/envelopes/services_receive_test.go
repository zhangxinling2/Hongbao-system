package envelopes

import (
	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	. "github.com/smartystreets/goconvey/convey"
	acservices "github.com/zhangxinling2/account/services"
	service "github.com/zhangxinling2/resk/services"
	"strconv"
	"testing"
)

func TestRedEnvelopeService_Receive(t *testing.T) {
	as := acservices.GetAccountService()
	Convey("收红包测试用例", t, func() {
		accounts := make([]*acservices.AccountDTO, 0, 10)
		size := 10
		for i := 0; i < size; i++ {
			account := acservices.AccountCreatedDTO{
				UserId:      ksuid.New().Next().String(),
				UserName:    "测试用户" + strconv.Itoa(i+1),
				Amount:      "2000",
				AccountName: "测试账户" + strconv.Itoa(i+1),
				AccountType: int(acservices.EnvelopeAccountType),
				CurrentCode: "CNY",
			}
			acDto, err := as.CreateAccount(account)
			So(err, ShouldBeNil)
			So(acDto, ShouldNotBeEmpty)
			accounts = append(accounts, acDto)
		}
		acDto := accounts[0]
		rs := service.GetRedEnvelopeService()
		Convey("收普通红包", func() {
			goods := service.RedEnvelopeSendingDTO{
				UserId:       acDto.UserId,
				UserName:     acDto.UserName,
				EnvelopeType: service.GeneralEnvelopeType,
				Amount:       decimal.NewFromFloat(1.88),
				Quantity:     size,
				Blessing:     service.DefaultBlessing,
			}
			activity, err := rs.SendOut(goods)
			So(err, ShouldBeNil)
			So(activity, ShouldNotBeNil)
			remainAmount := activity.RemainAmount
			for _, account := range accounts {
				rcv := service.RedEnvelopeReceiveDTO{
					EnvelopeNo:   activity.EnvelopeNo,
					RecvUserId:   account.UserId,
					RecvUsername: account.UserName,
					AccountNo:    account.AccountNo,
				}
				item, err := rs.Receive(rcv)
				So(err, ShouldBeNil)
				So(item.Amount.String(), ShouldEqual, activity.AmountOne.String())
				remainAmount = remainAmount.Sub(item.Amount)
				So(item.RemainAmount.String(), ShouldEqual, remainAmount.String())
			}
		})
		Convey("收拼运气红包", func() {
			goods := service.RedEnvelopeSendingDTO{
				UserId:       acDto.UserId,
				UserName:     acDto.UserName,
				EnvelopeType: service.LuckyEnvelopeType,
				Amount:       decimal.NewFromFloat(18.8),
				Quantity:     size,
				Blessing:     service.DefaultBlessing,
			}
			activity, err := rs.SendOut(goods)
			So(err, ShouldBeNil)
			So(activity, ShouldNotBeNil)
			total := decimal.NewFromFloat(0)
			for _, account := range accounts {
				rcv := service.RedEnvelopeReceiveDTO{
					EnvelopeNo:   activity.EnvelopeNo,
					RecvUserId:   account.UserId,
					RecvUsername: account.UserName,
					AccountNo:    account.AccountNo,
				}
				item, err := rs.Receive(rcv)
				So(err, ShouldBeNil)
				total = total.Add(item.Amount)
			}
			So(total.String(), ShouldEqual, goods.Amount.String())
		})
	})
}
func TestRedEnvelopeService_Receive_Failure(t *testing.T) {
	as := acservices.GetAccountService()
	Convey("收红包测试用例", t, func() {
		accounts := make([]*acservices.AccountDTO, 0, 10)
		size := 5
		for i := 0; i < size; i++ {
			account := acservices.AccountCreatedDTO{
				UserId:      ksuid.New().Next().String(),
				UserName:    "测试用户RF" + strconv.Itoa(i+1),
				Amount:      "30",
				AccountName: "测试账户RF" + strconv.Itoa(i+1),
				AccountType: int(acservices.EnvelopeAccountType),
				CurrentCode: "CNY",
			}
			acDto, err := as.CreateAccount(account)
			So(err, ShouldBeNil)
			So(acDto, ShouldNotBeEmpty)
			accounts = append(accounts, acDto)
		}
		acDto := accounts[0]
		rs := service.GetRedEnvelopeService()
		Convey("收普通红包失败", func() {
			goods := service.RedEnvelopeSendingDTO{
				UserId:       acDto.UserId,
				UserName:     acDto.UserName,
				EnvelopeType: service.GeneralEnvelopeType,
				Amount:       decimal.NewFromFloat(10),
				Quantity:     3,
				Blessing:     service.DefaultBlessing,
			}
			activity, err := rs.SendOut(goods)
			So(err, ShouldBeNil)
			So(activity, ShouldNotBeNil)
			remainAmount := activity.RemainAmount
			for i, account := range accounts {
				rcv := service.RedEnvelopeReceiveDTO{
					EnvelopeNo:   activity.EnvelopeNo,
					RecvUserId:   account.UserId,
					RecvUsername: account.UserName,
					AccountNo:    account.AccountNo,
				}
				item, err := rs.Receive(rcv)
				if i <= 2 {
					So(err, ShouldBeNil)
					So(item.Amount.String(), ShouldEqual, activity.AmountOne.String())
					remainAmount = remainAmount.Sub(item.Amount)
					So(item.RemainAmount.String(), ShouldEqual, remainAmount.String())
				} else {
					So(err, ShouldNotBeNil)
					So(item, ShouldBeNil)
				}

			}
		})
		Convey("收拼运气红包失败", func() {
			goods := service.RedEnvelopeSendingDTO{
				UserId:       acDto.UserId,
				UserName:     acDto.UserName,
				EnvelopeType: service.LuckyEnvelopeType,
				Amount:       decimal.NewFromFloat(30),
				Quantity:     3,
				Blessing:     service.DefaultBlessing,
			}
			activity, err := rs.SendOut(goods)
			So(err, ShouldBeNil)
			So(activity, ShouldNotBeNil)
			total := decimal.NewFromFloat(0)
			for i, account := range accounts {
				rcv := service.RedEnvelopeReceiveDTO{
					EnvelopeNo:   activity.EnvelopeNo,
					RecvUserId:   account.UserId,
					RecvUsername: account.UserName,
					AccountNo:    account.AccountNo,
				}
				item, err := rs.Receive(rcv)
				if i <= 2 {
					So(err, ShouldBeNil)
					total = total.Add(item.Amount)
				} else {
					So(err, ShouldNotBeNil)
					So(item, ShouldBeNil)
				}
			}
			So(total.String(), ShouldEqual, goods.Amount.String())
		})
	})
}
