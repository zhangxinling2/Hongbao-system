package base

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	irisrecover "github.com/kataras/iris/v12/middleware/recover"
	"github.com/sirupsen/logrus"
	"os"
	"resk/infra"
	"time"
)

var irisApplication *iris.Application

func Iris() *iris.Application {
	Check(irisApplication)
	return irisApplication
}

type IrisApplicationStarter struct {
	infra.BaseStarter
}

func (i *IrisApplicationStarter) Init(ctx infra.StarterContext) {
	app := initIris()
	logger := app.Logger()
	logger.Install(logrus.StandardLogger())
}

func (i *IrisApplicationStarter) Start(ctx infra.StarterContext) {
	app := Iris()
	routes := app.GetRoutes()
	for _, r := range routes {
		r.Trace(os.Stdout, -1)
	}

	port := ctx.Props().GetDefault("app.service.port", "18080")
	app.Run(iris.Addr(":" + port))
}
func (i *IrisApplicationStarter) StartBlocking() bool {
	return true
}
func initIris() *iris.Application {
	app := iris.New()
	app.Use(irisrecover.New())
	cfg := logger.Config{
		Status: true,
		IP:     true,
		Method: true,
		Path:   true,
		Query:  true,
		LogFunc: func(endTime time.Time, latency time.Duration, status, ip, method, path string, message interface{}, headerMessage interface{}) {
			app.Logger().Infof("| %s | %s | %s | %s | %s | %s | %s | %s |", endTime.Format("2006-01-02.15:04:05.000000"),
				latency.String(), status, ip, method, path, message, headerMessage)
		},
	}
	app.Use(logger.New(cfg))
	irisApplication = app
	return app
}
