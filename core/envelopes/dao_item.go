package envelopes

import (
	log "github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
)

type RedEnvelopeItemDao struct {
	runner *dbx.TxRunner
}

//查询 GetOne
func (dao *RedEnvelopeItemDao) GetOne(itemNo string) *RedEnvelopeItem {
	form := &RedEnvelopeItem{ItemNo: itemNo}
	ok, err := dao.runner.GetOne(form)
	if err != nil {
		return nil
	}
	if !ok {
		return nil
	}
	return form
}

//红包订单详情数据的写入 Insert
func (dao *RedEnvelopeItemDao) Insert(form *RedEnvelopeItem) (int64, error) {
	rs, err := dao.runner.Insert(form)
	if err != nil {
		return 0, err
	}
	return rs.LastInsertId()
}

func (dao *RedEnvelopeItemDao) FindItems(envelopeNo string) []*RedEnvelopeItem {
	var items []*RedEnvelopeItem
	sql := "select * from red_envelope_item where envelope_no = ?"
	err := dao.runner.Find(&items, sql, envelopeNo)
	if err != nil {
		log.Error(err)
		return nil
	}
	return items
}
