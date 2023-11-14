package main

import (
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
	"github.com/zhangxinling2/infra"
	_ "github.com/zhangxinling2/resk"
)

func main() {
	file := kvs.GetCurrentFilePath("config.ini", 1)
	conf := ini.NewIniFileConfigSource(file)
	app := infra.New(conf)
	app.Start()
}
