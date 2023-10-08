package base

import (
	"context"
	"github.com/tietang/dbx"
)

func TxContext(ctx context.Context, fn func(runner *dbx.TxRunner) error) error {
	return DbxDataBase().Tx(fn)
}
func Tx(fn func(runner *dbx.TxRunner) error) error {
	return TxContext(context.Background(), fn)
}
func WithValueContext(parent context.Context, runner *dbx.TxRunner) context.Context {
	return context.WithValue(parent, TX, runner)
}
