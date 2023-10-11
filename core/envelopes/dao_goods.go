package envelopes

import (
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
	services "resk/services"
	"time"
)

type EnvelopeDao struct {
	runner *dbx.TxRunner
}

func (dao *EnvelopeDao) Insert(po *RedEnvelopeGoods) (int64, error) {
	rs, err := dao.runner.Insert(po)
	if err != nil {
		return 0, err
	}
	return rs.LastInsertId()
}
func (dao *EnvelopeDao) GetOne(envelopeNo string) *RedEnvelopeGoods {
	good := &RedEnvelopeGoods{EnvelopeNo: envelopeNo}
	ok, err := dao.runner.GetOne(good)
	if err != nil || !ok {
		log.Error(err)
		return nil
	}
	return good
}
func (dao *EnvelopeDao) FindExpired(offset, size int) []RedEnvelopeGoods {
	var goods []RedEnvelopeGoods
	now := time.Now()
	sql := "select * from red_envelope_goods where order_type=1 and remain_quantity >0 and expired_at > ? and status in(1,2,3,6) limit ?,?"
	err := dao.runner.Find(&goods, sql, now, offset, size)
	if err != nil {
		log.Error(err)
	}
	return goods
}

//更新红包余额和数量
//不再使用事务行锁来更新红包余额和数量
//使用乐观锁来保证更新红包余额和数量的安全，避免负库存问题
//通过在where子句中判断红包剩余金额和数量来解决2个问题：
//1. 负库存问题，避免红包剩余金额和数量不够时进行扣减更新
//2. 减少实际的数据更新，也就是过滤掉部分无效的更新，提高总体性能
func (dao *EnvelopeDao) UpdateBalance(envelopeNo string, amount decimal.Decimal) (int64, error) {
	sql := "update red_envelope_goods set remain_amount = remain_amount-CAST(? AS DECIMAL(30,6)),remain_quantity=remain_quantity-1 where envelope_no=?" +
		" and remain_amount >=CAST(? AS DECIMAL(30,6)) and remain_quantity > 0"
	rs, err := dao.runner.Exec(sql, amount.String(), envelopeNo, amount.String())
	if err != nil {
		return 0, err
	}
	return rs.RowsAffected()
}
func (dao *EnvelopeDao) UpdateOrderStatus(envelopeNo string, status services.OrderStatus) (int64, error) {
	sql := " update red_envelope_goods" +
		" set status=? " +
		" where envelope_no=?"
	rs, err := dao.runner.Exec(sql, status, envelopeNo)
	if err != nil {
		return 0, err
	}
	return rs.RowsAffected()
}
