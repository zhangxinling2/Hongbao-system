package testx

import (
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"

	"github.com/zhangxinling2/infra"
	"github.com/zhangxinling2/infra/base"
)

func init() {
	file := kvs.GetCurrentFilePath("../brun/config.ini", 1)
	conf := ini.NewIniFileConfigSource(file)
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxDataBaseStarter{})
	infra.Register(&base.ValidatorStarter{})
	app := infra.New(conf)
	app.Start()
}
