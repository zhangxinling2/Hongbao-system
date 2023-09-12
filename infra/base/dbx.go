package base

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
	"github.com/tietang/props/kvs"
	"resk/infra"
)

var database = &dbx.Database{}

func DbxDataBase() *dbx.Database {
	return database
}

type DbxDataBaseStarter struct {
	infra.BaseStarter
}

func (s DbxDataBaseStarter) Setup(ctx infra.StarterContext) {
	conf := ctx.Props()
	setting := dbx.Settings{}
	err := kvs.Unmarshal(conf, &setting, "mysql")
	if err != nil {
		panic(err)
	}
	logrus.Infof("%+v\n", setting)
	logrus.Info("mysql.conn url:", setting.ShortDataSourceName())
	dbx, err := dbx.Open(setting)
	if err != nil {
		panic(err)
	}
	database = dbx
}
