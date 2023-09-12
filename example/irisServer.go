package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
	"strconv"
)

func main() {
	app := iris.Default()

	app.Get("/hello", func(context iris.Context) {
		context.WriteString("hello world")
	})
	v1 := app.Party("/v1")
	v1.Use(func(ctx iris.Context) {
		logrus.Info("自定义中间件")
		ctx.Next()
	})
	v1.Get("/user/{id:uint64 min(2)}", func(context iris.Context) {
		id := context.Params().GetUint64Default("id", 0)
		context.WriteString(strconv.Itoa(int(id)))
	})
	v1.Get("/orders/{action:string prefix(a_)}", func(ctx iris.Context) {
		a := ctx.Params().Get("action")
		ctx.WriteString(a)
	})
	app.OnAnyErrorCode(func(context iris.Context) {
		context.WriteString("出现错误")
	})
	err := app.Run(iris.Addr(":8082"))
	fmt.Println(err)
}
