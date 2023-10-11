package envelopes

import (
	"github.com/tietang/dbx"
	"resk/infra/base"
)

type ExpiredEnvelopeDomain struct {
}

func (e *ExpiredEnvelopeDomain) Next() (ok bool) {
	err := base.Tx(func(runner *dbx.TxRunner) error {

	})
}