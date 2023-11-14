package envelopes

import (
	"context"
	"github.com/segmentio/ksuid"
	"github.com/tietang/dbx"
	"github.com/zhangxinling2/infra/base"
	services "github.com/zhangxinling2/resk/services"
)

type itemDomain struct {
	RedEnvelopeItem
}

func (i *itemDomain) createItemNo() {
	i.ItemNo = ksuid.New().Next().String()
}
func (i *itemDomain) Create(item services.RedEnvelopeItemDTO) {
	i.RedEnvelopeItem.FromDTO(&item)
	i.createItemNo()
}
func (i *itemDomain) Save(ctx context.Context) (id int64, err error) {
	err = base.ExecuteContext(ctx, func(runner *dbx.TxRunner) error {
		dao := RedEnvelopeItemDao{runner: runner}
		id, err = dao.Insert(&i.RedEnvelopeItem)
		return err
	})
	return id, err
}
func (i *itemDomain) GetOne(ctx context.Context, itemNo string) (dto *services.RedEnvelopeItemDTO) {
	err := base.ExecuteContext(ctx, func(runner *dbx.TxRunner) error {
		dao := RedEnvelopeItemDao{runner: runner}
		item := dao.GetOne(itemNo)
		if item != nil {
			dto = item.ToDTO()
		}
		return nil
	})
	if err != nil {
		return nil
	}
	return dto
}
func (i *itemDomain) FindItems(envelopeNo string) (itemDTOs []*services.RedEnvelopeItemDTO) {
	var items []*RedEnvelopeItem
	err := base.Tx(func(runner *dbx.TxRunner) error {
		dao := RedEnvelopeItemDao{runner: runner}
		items = dao.FindItems(envelopeNo)
		return nil
	})
	if err != nil {
		return nil
	}
	itemDTOs = make([]*services.RedEnvelopeItemDTO, 0)
	for _, po := range items {
		item := po.ToDTO()
		itemDTOs = append(itemDTOs, item)
	}
	return itemDTOs
}
