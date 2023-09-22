package accounts

import (
	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	. "github.com/smartystreets/goconvey/convey"
	services "resk/services"
	"testing"
)

func TestAccountDomain_Create(t *testing.T) {
	accountDTO := services.AccountDTO{
		AccountCreatedDTO: services.AccountCreatedDTO{
			UserId:   ksuid.New().Next().String(),
			UserName: "测试用户",
		},
		Balance: decimal.NewFromFloat(100),
		Status:  1,
	}
	domain := new(accountDomain)
	Convey("用户创建", t, func() {
		rs, err := domain.Create(accountDTO)
		So(err, ShouldBeNil)
		So(rs, ShouldNotBeNil)
		So(rs.Balance.String(), ShouldEqual, accountDTO.Balance.String())
		So(rs.UserId, ShouldEqual, accountDTO.UserId)
		So(rs.UserName, ShouldEqual, accountDTO.UserName)
		So(rs.Status, ShouldEqual, accountDTO.Status)
	})
}
func TestAccountDomain_Transfer(t *testing.T) {
	adto1 := &services.AccountDTO{
		AccountCreatedDTO: services.AccountCreatedDTO{
			UserId:      ksuid.New().Next().String(),
			AccountType: int(services.EnvelopeAccountType),
			UserName:    "测试用户1",
		},
		Balance: decimal.NewFromFloat(100),
		Status:  1,
	}
	adto2 := &services.AccountDTO{
		AccountCreatedDTO: services.AccountCreatedDTO{
			UserId:      ksuid.New().Next().String(),
			AccountType: int(services.EnvelopeAccountType),
			UserName:    "测试用户2",
		},
		Balance: decimal.NewFromFloat(100),
		Status:  1,
	}
	domain := accountDomain{}
	Convey("Transfer测试", t, func() {
		acc1, err := domain.Create(*adto1)
		So(err, ShouldBeNil)
		So(adto1, ShouldNotBeNil)
		So(adto1.Balance.String(), ShouldEqual, acc1.Balance.String())
		So(adto1.UserId, ShouldEqual, acc1.UserId)
		So(adto1.UserName, ShouldEqual, acc1.UserName)
		So(adto1.Status, ShouldEqual, acc1.Status)
		adto1 = acc1

		acc2, err := domain.Create(*adto2)
		So(err, ShouldBeNil)
		So(adto2, ShouldNotBeNil)
		So(adto2.Balance.String(), ShouldEqual, acc2.Balance.String())
		So(adto2.UserId, ShouldEqual, acc2.UserId)
		So(adto2.UserName, ShouldEqual, acc2.UserName)
		So(adto2.Status, ShouldEqual, acc2.Status)
		adto2 = acc2

		Convey("金额充足，转账", func() {
			amount := decimal.NewFromFloat(1)
			b1 := services.TradeParticipator{
				AccountNo: adto1.AccountNo,
				UserId:    adto1.UserId,
				UserName:  adto1.UserName,
			}
			b2 := services.TradeParticipator{
				AccountNo: adto2.AccountNo,
				UserId:    adto2.UserId,
				UserName:  adto2.UserName,
			}
			trade := services.AccountTransferDTO{
				TradeNo:     ksuid.New().Next().String(),
				TradeBody:   b1,
				TradeTarget: b2,
				Amount:      amount,
				ChangeType:  services.ChangeType(-1),
				ChangeFlag:  services.FlagTransferOut,
				Decs:        "转账",
			}

			status, err := domain.Transfer(trade)
			So(err, ShouldBeNil)
			So(status, ShouldEqual, services.TransferSuccess)
			a2 := domain.GetAccount(adto1.AccountNo)
			So(a2, ShouldNotBeNil)
			So(a2.Balance.String(), ShouldEqual, adto1.Balance.Sub(amount).String())

		})

		Convey("余额不足，转账", func() {
			amount := adto1.Balance
			amount = amount.Add(decimal.NewFromFloat(200))
			b1 := services.TradeParticipator{
				AccountNo: adto1.AccountNo,
				UserId:    adto1.UserId,
				UserName:  adto1.UserName,
			}
			b2 := services.TradeParticipator{
				AccountNo: adto2.AccountNo,
				UserId:    adto2.UserId,
				UserName:  adto2.UserName,
			}
			trade := services.AccountTransferDTO{
				TradeNo:     ksuid.New().Next().String(),
				TradeBody:   b1,
				TradeTarget: b2,
				Amount:      amount,
				ChangeType:  services.ChangeType(-1),
				ChangeFlag:  services.FlagTransferOut,
				Decs:        "转账",
			}
			status, err := domain.Transfer(trade)
			So(err, ShouldNotBeNil)
			So(status, ShouldEqual, services.TransferSufficientFunds)
			a2 := domain.GetAccount(adto1.AccountNo)
			So(a2, ShouldNotBeNil)
			So(a2.Balance.String(), ShouldEqual, adto1.Balance.String())
		})
		Convey("转入", func() {
			amount := decimal.NewFromFloat(1)
			b1 := services.TradeParticipator{
				AccountNo: adto1.AccountNo,
				UserId:    adto1.UserId,
				UserName:  adto1.UserName,
			}
			b2 := services.TradeParticipator{
				AccountNo: adto2.AccountNo,
				UserId:    adto2.UserId,
				UserName:  adto2.UserName,
			}
			trade := services.AccountTransferDTO{
				TradeNo:     ksuid.New().Next().String(),
				TradeBody:   b1,
				TradeTarget: b2,
				Amount:      amount,
				ChangeType:  services.AccountStoreValue,
				ChangeFlag:  services.FlagTransferIn,
				Decs:        "储值",
			}
			status, err := domain.Transfer(trade)
			So(err, ShouldBeNil)
			So(status, ShouldEqual, services.TransferSuccess)
			a2 := domain.GetAccount(adto1.AccountNo)
			So(a2, ShouldNotBeNil)
			So(a2.Balance.String(), ShouldEqual, adto1.Balance.Add(amount).String())
		})
	})
}
