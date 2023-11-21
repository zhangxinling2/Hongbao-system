package main

import (
	"flag"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/tietang/props/consul"
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
	"github.com/zhangxinling2/infra"
	"github.com/zhangxinling2/infra/base"
	_ "github.com/zhangxinling2/resk"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		log.Info(http.ListenAndServe(":6060", nil))
	}()
	flag.Parse()
	profile := flag.Arg(0)
	if profile == "" {
		profile = "dev"
	}
	file := kvs.GetCurrentFilePath("boot.ini", 1)
	conf := ini.NewIniFileConfigSource(file)
	addr := conf.GetDefault("consul.address", "127.0.0.1:8500")
	contexts := conf.KeyValue("consul.contexts").Strings()
	if _, e := conf.Get("profile"); e != nil {
		conf.Set("profile", profile)
	}
	consulConf := consul.NewConsulKeyValueCompositeConfigSource(contexts, addr)
	consulConf.Add(conf)
	base.InitLog(consulConf)
	app := infra.New(consulConf)
	app.Start()
}
