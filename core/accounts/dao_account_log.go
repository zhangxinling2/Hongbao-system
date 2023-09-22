package accounts

import (
	"github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
)

type AccountLogDAO struct {
	runner *dbx.TxRunner
}

func (dao *AccountLogDAO) GetOne(logNo string) *AccountLog {
	a := &AccountLog{LogNo: logNo}
	ok, err := dao.runner.GetOne(a)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	if !ok {
		return nil
	}
	return a
}
func (dao *AccountLogDAO) GetByTradeNo(tradeNo string) *AccountLog {
	sql := "select * from account_log where trade_no=?"
	a := &AccountLog{TradeNo: tradeNo}
	ok, err := dao.runner.Get(a, sql, tradeNo)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	if !ok {
		return nil
	}
	return a
}
func (dao *AccountLogDAO) Insert(l *AccountLog) (id int64, err error) {
	rs, err := dao.runner.Insert(l)
	if err != nil {
		return 0, err
	}
	return rs.LastInsertId()
}
