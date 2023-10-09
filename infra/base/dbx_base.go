package base

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
)

const TX = "tx"

func TxContext(ctx context.Context, fn func(runner *dbx.TxRunner) error) error {
	return DbxDataBase().Tx(fn)
}
func Tx(fn func(runner *dbx.TxRunner) error) error {
	return TxContext(context.Background(), fn)
}
func WithValueContext(parent context.Context, runner *dbx.TxRunner) context.Context {
	return context.WithValue(parent, TX, runner)
}
func ExecuteContext(ctx context.Context, fn func(runner *dbx.TxRunner) error) error {
	tx, ok := ctx.Value(TX).(*dbx.TxRunner)
	if !ok || tx == nil {
		log.Panic("是否在事务函数中使用?")
	}
	return fn(tx)
}
