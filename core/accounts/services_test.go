package accounts

import (
	"github.com/segmentio/ksuid"
	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	services "resk/services"
	"testing"
)

func TestAccountService_CreateAccount(t *testing.T) {
	accountDTO := services.AccountCreatedDTO{
		UserId:      ksuid.New().Next().String(),
		UserName:    "测试用户",
		Amount:      "100",
		AccountName: "测试账户",
		AccountType: 2,
		CurrentCode: "CNY",
	}
	service := new(AccountService)
	Convey("用户创建", t, func() {
		rs, err := service.CreateAccount(accountDTO)
		So(err, ShouldBeNil)
		So(rs, ShouldNotBeNil)
		So(rs.Balance.String(), ShouldEqual, accountDTO.Amount)
		So(rs.UserId, ShouldEqual, accountDTO.UserId)
		So(rs.UserName, ShouldEqual, accountDTO.UserName)
		So(rs.Status, ShouldEqual, 1)
	})
}

func TestAccountService_Transfer(t *testing.T) {
	a1 := services.AccountCreatedDTO{
		UserId:      ksuid.New().Next().String(),
		UserName:    "测试用户1",
		AccountName: "测试用户1",
		AccountType: 2,
		CurrentCode: "CNY",
		Amount:      "100",
	}
	a2 := services.AccountCreatedDTO{
		UserId:      ksuid.New().Next().String(),
		UserName:    "测试用户2",
		AccountName: "测试用户2",
		AccountType: 2,
		CurrentCode: "CNY",
		Amount:      "100",
	}
	service := AccountService{}
	dto1, err := service.CreateAccount(a1)
	if err != nil {
		logrus.Error(err)
	}
	dto2, err := service.CreateAccount(a2)
	if err != nil {
		logrus.Error(err)
	}
	p1 := services.TradeParticipator{
		AccountNo: dto1.AccountNo,
		UserId:    dto1.UserId,
		UserName:  dto1.UserName,
	}
	p2 := services.TradeParticipator{
		AccountNo: dto2.AccountNo,
		UserId:    dto2.UserId,
		UserName:  dto2.UserName,
	}
	Convey("测试服务", t, func() {
		Convey("转账，余额充足", func() {
			transfer := services.AccountTransferDTO{
				TradeNo:     ksuid.New().Next().String(),
				TradeBody:   p1,
				TradeTarget: p2,
				AmountStr:   "",
				Amount:      decimal.Decimal{},
				ChangeType:  0,
				ChangeFlag:  0,
				Decs:        "",
			}
		})
	})
}
